package registry

import (
	//"context"
	"database/sql"
	"github.com/gin/demo/internal/application"
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

	return registry
}
