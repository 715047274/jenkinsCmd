package registry

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

// var Module = fx.Option()

func runApplication(lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("OnStart..........")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("OnStop.........")
			return nil
		},
	})
}
