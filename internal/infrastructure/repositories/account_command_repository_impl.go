package repositories

import (
	"database/sql"
)

type accountCommandRepositoryImpl struct {
	db *sql.DB
}

func NewAccountCommandRepository(db *sql.DB) AccountCommandRepository {
	return &accountCommandRepositoryImpl{db: db}
}
func (a *accountCommandRepositoryImpl) UpdateAccount(userName string, balance int64, currency string) error {
	//TODO implement me
	// panic("implement me")
	query := "UPDATE accounts SET balance = ? where username = ? and currency = ?"
	_, err := a.db.Exec(query, userName, balance, currency)
	return err
}

func (a *accountCommandRepositoryImpl) RemoveAccount(userName string, currency string) error {
	//TODO implement me
	// panic("implement me")
	query := "delete from accounts where username = ? and currency =?"
	_, err := a.db.Exec(query, userName, currency)
	return err
}

func (a *accountCommandRepositoryImpl) CreateAccount(userName string, balance int64, currency string) error {
	// panic("implement me")

	query := "INSERT INTO accounts (owner, balance, currency) VALUES (?, ?,? ) "
	_, err := a.db.Exec(query, userName, balance, currency)

	return err
}
