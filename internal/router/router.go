package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/internal/application"
	_ "github.com/gin/demo/internal/db"
	_ "github.com/gin/demo/internal/infrastructure/repositories"
	"github.com/gin/demo/internal/registry"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, registry *registry.ServiceRegistry) {

	// Retrieve account service dynamically from the registry
	accountServiceInterface, err := registry.GetService("AccountService")
	if err != nil {
		panic("AccountService not found in registry")
	}
	accountService, ok := accountServiceInterface.(application.AccountService)
	if !ok {
		panic("invlid account service type")
	}

	engine.GET("/accounts", func(c *gin.Context) {
		items, err := accountService.GetAccount("kevin")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, items)
		}
	})

	engine.POST("/accounts/add", func(c *gin.Context) {
		err := accountService.AddAccount("kevintest", 102, "USD")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "item added"})
		}
	})

}
