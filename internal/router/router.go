package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/internal/application"
	_ "github.com/gin/demo/internal/db"
	_ "github.com/gin/demo/internal/infrastructure/repositories"
	"github.com/gin/demo/internal/registry"
	"net/http"
)

//func ApiRoutes(prefix string, router *gin.Engine) {
//	db := database.ConnectDb()
//	apiGroup := router.Group(prefix)
//	{
//		dashboard := apiGroup.Group("/dashboard/account")
//		{
//			accountRepo := repositories.NewAccountRepository(db)
//			accountService := service.NewAccountService(accountRepo)
//			accountController := controller.NewAccountController(accountService)
//
//			dashboard.GET("/all", accountController.GetAccountList)
//		}
//	}
//
//}
//
//type Route struct {
//	Gin *gin.Engine
//}

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

}
