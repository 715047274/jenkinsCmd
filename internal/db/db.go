package db

import (
	"database/sql"
	"log"
)

func ConnectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}

	return db
}
