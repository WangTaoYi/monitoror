//+build faker

package pingdom

import (
	uiConfig "github.com/monitoror/monitoror/api/config/usecase"
	coreConfig "github.com/monitoror/monitoror/config"
	coreModels "github.com/monitoror/monitoror/models"
	"github.com/monitoror/monitoror/monitorables/pingdom/api"
	pingdomDelivery "github.com/monitoror/monitoror/monitorables/pingdom/api/delivery/http"
	pingdomModels "github.com/monitoror/monitoror/monitorables/pingdom/api/models"
	pingdomUsecase "github.com/monitoror/monitoror/monitorables/pingdom/api/usecase"
	"github.com/monitoror/monitoror/service/store"
)

type Monitorable struct {
	store *store.Store
}

func NewMonitorable(store *store.Store) *Monitorable {
	monitorable := &Monitorable{}
	monitorable.store = store

	// Register Monitorable Tile in config manager
	store.UIConfigManager.RegisterTile(api.PingdomCheckTileType, monitorable.GetVariants(), uiConfig.MinimalVersion)

	return monitorable
}

func (m *Monitorable) GetDisplayName() string { return "Pingdom (faker)" }
func (m *Monitorable) GetVariants() []string {
	return []coreModels.VariantName{coreConfig.DefaultVariant}
}
func (m *Monitorable) Validate(variant string) (bool, error) { return true, nil }

func (m *Monitorable) Enable(variant string) {
	usecase := pingdomUsecase.NewPingdomUsecase()
	delivery := pingdomDelivery.NewPingdomDelivery(usecase)

	// EnableTile route to echo
	route := m.store.MonitorableRouter.RouterGroup("/pingdom", variant).GET("/pingdom", delivery.GetCheck)

	// EnableTile data for config hydration
	m.store.UIConfigManager.EnableTile(api.PingdomCheckTileType, variant,
		&pingdomModels.ChecksParams{}, route.Path, coreConfig.DefaultInitialMaxDelay)
}
