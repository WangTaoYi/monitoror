//+build faker

package port

import (
	uiConfig "github.com/monitoror/monitoror/api/config/usecase"
	coreConfig "github.com/monitoror/monitoror/config"
	coreModels "github.com/monitoror/monitoror/models"
	"github.com/monitoror/monitoror/monitorables/port/api"
	portDelivery "github.com/monitoror/monitoror/monitorables/port/api/delivery/http"
	portModels "github.com/monitoror/monitoror/monitorables/port/api/models"
	portUsecase "github.com/monitoror/monitoror/monitorables/port/api/usecase"
	"github.com/monitoror/monitoror/service/store"
)

type Monitorable struct {
	store *store.Store
}

func NewMonitorable(store *store.Store) *Monitorable {
	monitorable := &Monitorable{}
	monitorable.store = store

	// Register Monitorable Tile in config manager
	store.UIConfigManager.RegisterTile(api.PortTileType, monitorable.GetVariants(), uiConfig.MinimalVersion)

	return monitorable
}

func (m *Monitorable) GetDisplayName() string { return "Port (faker)" }
func (m *Monitorable) GetVariants() []string {
	return []coreModels.VariantName{coreConfig.DefaultVariant}
}
func (m *Monitorable) Validate(variant string) (bool, error) { return true, nil }

func (m *Monitorable) Enable(variant string) {
	usecase := portUsecase.NewPortUsecase()
	delivery := portDelivery.NewPortDelivery(usecase)

	// EnableTile route to echo
	route := m.store.MonitorableRouter.RouterGroup("/port", variant).GET("/port", delivery.GetPort)

	// EnableTile data for config hydration
	m.store.UIConfigManager.EnableTile(api.PortTileType, variant,
		&portModels.PortParams{}, route.Path, coreConfig.DefaultInitialMaxDelay)
}
