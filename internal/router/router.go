package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/internal/controller"
	database "github.com/gin/demo/internal/db"
	"github.com/gin/demo/internal/infrastructure/repositories"
	"github.com/gin/demo/internal/registry"
	"github.com/gin/demo/internal/service"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	db := database.ConnectDb()
	apiGroup := router.Group(prefix)
	{
		dashboard := apiGroup.Group("/dashboard/account")
		{
			accountRepo := repositories.NewAccountRepository(db)
			accountService := service.NewAccountService(accountRepo)
			accountController := controller.NewAccountController(accountService)

			dashboard.GET("/all", accountController.GetAccountList)
		}
	}

}

type Route struct {
	Gin *gin.Engine
}

func RegisterRoutes(engine *gin.Engine, registry *registry.ServiceRegistry) {

}
