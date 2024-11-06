package service

import (
	"github.com/gin/demo/internal/model"
	"github.com/gin/demo/internal/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s *AccountService) GetAllItems() (*[]model.Account, error) {
	accountList, err := s.accountRepository.GetAllItems()
	return accountList, err
}
