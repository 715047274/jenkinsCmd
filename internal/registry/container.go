package registry

import (
	//"context"
	"database/sql"
	"github.com/gin/demo/internal/application"
	"github.com/gin/demo/internal/infrastructure/repositories"
	//"fmt"
	//"go.uber.org/fx"
)

//// var Module = fx.Option()
//
//func runApplication(lifecycle fx.Lifecycle) {
//	lifecycle.Append(fx.Hook{
//		OnStart: func(ctx context.Context) error {
//			fmt.Println("OnStart..........")
//			return nil
//		},
//		OnStop: func(ctx context.Context) error {
//			fmt.Println("OnStop.........")
//			return nil
//		},
//	})
//}

func NewContainer(db *sql.DB) *ServiceRegistry {
	// Initialize the service registry
	registry := NewServiceRegistry()

	// Register repositories
	dispatcher := application.NewEventDispatcher()

	//register repositories as services
	registry.RegisterService("EventDispatcher", dispatcher)

	accountQueryRepo := repositories.NewAccountQueryRepository(db)
	accountCommRepo := repositories.NewAccountCommandRepository(db)

	registry.RegisterService("AccountQueryRepository", accountQueryRepo)
	registry.RegisterService("AccountCommandRepository", accountCommRepo)

	// create abd register the account service dynamically using repositories
	accountService := application.NewAccountService(accountQueryRepo, accountCommRepo, dispatcher)
	inventoryService := &application.InventoryService{}
	analysisService := &application.AnalyticService{}

	registry.RegisterService("AccountService", accountService)

	dispatcher.Subscribe("AccountCreate", accountService.HandleAccountCreate)
	dispatcher.Subscribe("AccountCreate", inventoryService.HandleCarEvent)
	dispatcher.Subscribe("AccountCreate", analysisService.HandleCarEvent)

	return registry
}
