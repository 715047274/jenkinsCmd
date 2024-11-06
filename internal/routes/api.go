package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/internal/controller"
	database "github.com/gin/demo/internal/db"
	"github.com/gin/demo/internal/repository"
	"github.com/gin/demo/internal/service"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	db := database.ConnectDb()
	apiGroup := router.Group(prefix)
	{
		dashboard := apiGroup.Group("/dashboard/account")
		{
			accountRepo := repository.NewAccountRepository(db)
			accountService := service.NewAccountService(accountRepo)
			accountController := controller.NewAccountController(accountService)

			dashboard.GET("/all", accountController.GetAccountList)
		}
	}

}

type Route struct {
	Gin *gin.Engine
}

//
