package domain

import "time"

type AccountItem struct {
	Id       int       `db:"id" json:"id"`
	Owner    string    `db:"owner" json:"owner"`
	Currency string    `db:"currency" json:"currency"`
	Balance  int16     `db:"balance" json:"balance"`
	Created  time.Time `db:"created_at" json:"created"'`
}
