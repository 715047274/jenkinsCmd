package repositories

import (
	"database/sql"
)

type AccountCommandRepositoryImpl struct {
	db *sql.DB
}
