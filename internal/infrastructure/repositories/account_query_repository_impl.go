package repositories

import (
	"database/sql"
	"fmt"
	"github.com/gin/demo/internal/domain"
)

type accountQueryRepositoryImpl struct {
	db *sql.DB
}

func (r *accountQueryRepositoryImpl) GetAllAccountBalance(currency string) ([]domain.AccountItem, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountQueryRepository(db *sql.DB) AccountQueryRepository {
	return &accountQueryRepositoryImpl{db: db}
}

func (r *accountQueryRepositoryImpl) GetAccountBalance(userName string) ([]domain.AccountItem, error) {
	rows, err := r.db.Query("SELECT * FROM Accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accountList []domain.AccountItem
	for rows.Next() {
		var account domain.AccountItem
		if err = rows.Scan(&account.Id, &account.Owner, &account.Balance, &account.Currency, &account.Created); err != nil {
			fmt.Println("Error query scan - %s", err)
			return nil, err
		}
		accountList = append(accountList, account)
	}
	return accountList, nil
}
