package repositories

import "github.com/gin/demo/internal/domain"

type AccountQueryRepository interface {
	GetAccountBalance(userName string) ([]domain.AccountItem, error)
	GetAllAccountBalance(currency string) ([]domain.AccountItem, error)
}
