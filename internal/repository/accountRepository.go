package repository

import (
	"database/sql"
	"fmt"
	"github.com/gin/demo/internal/model"
)

type AccountRepository struct {
	dbClient *sql.DB
}

func NewAccountRepository(dbClient *sql.DB) *AccountRepository {
	return &AccountRepository{dbClient: dbClient}
}

type AccountRepositoryInterface interface {
	// GetItemById(ID int) (*model.Account, error)
	GetAllItems() (*[]model.Account, error)
}

func (a *AccountRepository) GetAllItems() (*[]model.Account, error) {
	row, err := a.dbClient.Query("SELECT * FROM Accounts")
	if err != nil {
		fmt.Printf("Error select query - %s", err)
		return nil, err
	}
	var accountList []model.Account
	for row.Next() {
		var account model.Account
		err = row.Scan(&account.Id, &account.Owner, &account.Balance, &account.Currency, &account.Created)
		fmt.Println(&row)
		if err != nil {
			fmt.Println("Error query scan - %s", err)
			return nil, err
		}
		accountList = append(accountList, account)
	}
	return &accountList, nil
}
