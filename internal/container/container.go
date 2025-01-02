package container

import (
	"github.com/715047274/jenkinsCmd/internal/adapters"
	"github.com/715047274/jenkinsCmd/internal/config"
	"github.com/715047274/jenkinsCmd/internal/domain"
	"github.com/715047274/jenkinsCmd/internal/handlers"
	"github.com/715047274/jenkinsCmd/internal/services"
)

// AppContainer holds all application dependencies
type AppContainer struct {
	Config  config.Config
	Adapter *adapters.JenkinsAdapter
	Service domain.JobService
	Handler *handlers.JobHandler
}

// NewAppContainer creates and initializes all dependencies
func NewAppContainer(cfg config.Config) *AppContainer {
	adapter := adapters.NewJenkinsAdapter(cfg)
	service := services.NewJobService(adapter)
	handler := handlers.NewJobHandler(service)

	return &AppContainer{
		Config:  cfg,
		Adapter: adapter,
		Service: service,
		Handler: handler,
	}
}
