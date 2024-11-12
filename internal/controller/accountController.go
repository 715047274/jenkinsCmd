package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin/demo/internal/domain"
	"github.com/gin/demo/internal/service"
	"net/http"
)

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (h *AccountController) GetAccountList(c *gin.Context) {
	var accountList *[]domain.Account
	var err error
	accountList, err = h.accountService.GetAllItems()
	if err != nil {
		fmt.Printf("Error - %s", err)
	}
	c.JSON(http.StatusOK, accountList)
}
