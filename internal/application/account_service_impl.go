package application

import (
	"fmt"
	"github.com/gin/demo/internal/domain"
	"github.com/gin/demo/internal/infrastructure/repositories"
)

type accountServiceImpl struct {
	queryRepo   repositories.AccountQueryRepository
	commandRepo repositories.AccountCommandRepository
	dispatcher  EventDispatcher
}

func NewAccountService(
	queryRepo repositories.AccountQueryRepository,
	commandRepo repositories.AccountCommandRepository,
	dispatcher EventDispatcher,
) AccountService {
	return &accountServiceImpl{
		queryRepo:   queryRepo,
		commandRepo: commandRepo,
		dispatcher:  dispatcher,
	}
}

func (a *accountServiceImpl) AddAccount(userName string, balance int64, currency string) error {
	//TODO implement me
	//panic("implement me")

	err := a.commandRepo.CreateAccount(userName, balance, currency)
	if err != nil {
		return err
	}
	event := domain.AccountEvent{userName, balance, currency}
	a.dispatcher.Dispatch("AccountCreate", event)
	return nil
}

func (a *accountServiceImpl) GetAccount(userName string) ([]domain.AccountItem, error) {
	return a.queryRepo.GetAccountBalance(userName)
}

func (a *accountServiceImpl) HandleAccountCreate(event interface{}) {
	if accountEvent, ok := event.(domain.AccountEvent); ok {
		fmt.Printf("Running account event handle for UserID=%s, ProductID=%d, Quantity=%s\n",
			accountEvent.Owner, accountEvent.Balance, accountEvent.Currency)
	}
}
