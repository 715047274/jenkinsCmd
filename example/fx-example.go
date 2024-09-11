package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type Sample struct{}

func (sample *Sample) PrintSomething() {
	fmt.Println("print something")
}

func NewSample() *Sample {
	return &Sample{}
}

func RegisterToStart(
	lifeCycle fx.Lifecycle,
	sample *Sample,
) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			fmt.Println("start triggered")
			sample.PrintSomething()
			return nil
		},
		OnStop: func(_ context.Context) error {
			fmt.Println("stop server")
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(NewSample),
		fx.Invoke(RegisterToStart),
	)
	app.Run()
}
