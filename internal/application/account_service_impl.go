package application

import (
	"github.com/gin/demo/internal/domain"
	"github.com/gin/demo/internal/infrastructure/repositories"
)

type accountServiceImpl struct {
	queryRepo repositories.AccountQueryRepository
}

func NewAccountService(queryRepo repositories.AccountQueryRepository) AccountService {
	return &accountServiceImpl{queryRepo: queryRepo}
}

func (a accountServiceImpl) AddAccount(userName string, balance int64, currency string) {
	//TODO implement me
	panic("implement me")
}

func (a accountServiceImpl) GetAccount(userName string) ([]domain.AccountItem, error) {
	return a.queryRepo.GetAccountBalance(userName)
}
