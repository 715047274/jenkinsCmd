package service

import (
	"github.com/gin/demo/internal/domain"
	"github.com/gin/demo/internal/infrastructure/repositories"
)

type AccountService struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountService(accountRepository *repositories.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (s *AccountService) GetAllItems() (*[]domain.AccountItem, error) {
	accountList, err := s.accountRepository.GetAllItems()
	return accountList, err
}
