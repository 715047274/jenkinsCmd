package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/src/inter/controller"
	database "github.com/gin/demo/src/inter/db"
	"github.com/gin/demo/src/inter/repository"
	"github.com/gin/demo/src/inter/service"
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
