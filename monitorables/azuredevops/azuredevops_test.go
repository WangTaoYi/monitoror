package azuredevops

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"

	coreModels "github.com/monitoror/monitoror/models"

	coreConfigMocks "github.com/monitoror/monitoror/api/config/mocks"
	coreConfig "github.com/monitoror/monitoror/config"
	serviceMocks "github.com/monitoror/monitoror/service/mocks"
	"github.com/monitoror/monitoror/service/store"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
)

func TestNewMonitorable(t *testing.T) {
	// init Store
	mockRouterGroup := new(serviceMocks.MonitorableRouterGroup)
	mockRouterGroup.On("GET", Anything, Anything).Return(&echo.Route{Path: "/path"})
	mockRouter := new(serviceMocks.MonitorableRouter)
	mockRouter.On("RouterGroup", Anything, Anything).Return(mockRouterGroup)
	mockConfigManager := new(coreConfigMocks.Manager)
	mockConfigManager.On("RegisterTile", Anything, Anything, Anything)

	s := &store.Store{
		CoreConfig:        &coreConfig.Config{},
		MonitorableRouter: mockRouter,
		UIConfigManager:   mockConfigManager,
	}

	// init Env
	// OK
	_ = os.Setenv("MO_MONITORABLE_AZUREDEVOPS_URL", "https://azure.example.com/myProject1")
	_ = os.Setenv("MO_MONITORABLE_AZUREDEVOPS_TOKEN", "xxx")
	// Missing Token
	_ = os.Setenv("MO_MONITORABLE_AZUREDEVOPS_VARIANT1_URL", "https://azure.example.com/myProject2")
	// Url broken
	_ = os.Setenv("MO_MONITORABLE_AZUREDEVOPS_VARIANT2_URL", "url%sazure.example.com/myProject2")

	// NewMonitorable
	monitorable := NewMonitorable(s)
	assert.NotNil(t, monitorable)

	// GetDisplayName
	assert.NotNil(t, monitorable.GetDisplayName())

	// GetVariants
	assert.Len(t, monitorable.GetVariants(), 3)

	// Validate
	validate, err := monitorable.Validate(coreModels.DefaultVariant)
	assert.NoError(t, err)
	assert.True(t, validate)
	validate, err = monitorable.Validate("variant1")
	assert.Error(t, err)
	assert.False(t, validate)
	validate, err = monitorable.Validate("variant2")
	assert.Error(t, err)
	assert.False(t, validate)

	// Enable
	monitorable.Enable(coreModels.DefaultVariant)
}
