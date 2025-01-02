package main

import (
	"github.com/715047274/jenkinsCmd/internal/config"
	"github.com/715047274/jenkinsCmd/internal/container"
	"github.com/715047274/jenkinsCmd/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewEnvConfig("development")
	log.Println(cfg.GetJenkinsURL())

	//jenkinsAdapter := adapters.NewJenkinsAdapter(cfg)
	appContainer := container.NewAppContainer(cfg)
	// Initialize the Gin router
	r := gin.Default()
	r.Use(middleware.ConfigMiddleware(cfg))
	// Setup routes
	// routes.SetupJenkinsRoutes(r)
	appContainer.Handler.RegisterRoutes(r)
	// Start the server
	r.Run(":8282")
}
