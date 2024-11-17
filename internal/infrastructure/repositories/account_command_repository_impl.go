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
	panic("implement me")
}

func (a *accountCommandRepositoryImpl) RemoveAccount(userName string) error {
	//TODO implement me
	panic("implement me")
}

func (a *accountCommandRepositoryImpl) CreateAccount(userName string, balance int64, currency string) error {
	panic("implement me")
}
