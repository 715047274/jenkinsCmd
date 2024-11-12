package application

import "github.com/gin/demo/internal/domain"

type AccountService interface {
	AddAccount(userName string, balance int64, currency string)
	GetAccount(userName string) ([]domain.AccountItem, error)
}
