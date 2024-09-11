package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"log"

	//"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var DB *sql.DB

//var migrations embed.FS

func init() {

}

func runMigrations(db *sql.DB) {
	// Get database driver for SQLite
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Could not create SQLite driver: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance("file://db/migration",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}
	//
	//// Run migration UP
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Could not apply migrations: %v", err)
	} else if err == migrate.ErrNoChange {
		log.Println("No new migrations to apply")
	} else {
		log.Println("Migrations applied successfully")
	}
}
func main() {

	r := gin.Default()

	// Database connection
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("Failed to connect to the SQLite database: %v", err)
	}
	defer db.Close()

	// Run migrations
	runMigrations(db)

	// Define your routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server
	r.Run(":8080")
	fmt.Println("hello world")
}

type Account struct {
	Id       int    `db:"id" json:"id"`
	Owner    string `db:"owner" json:"owner"`
	Currency string `db:"currency" json:"currency"`
	Balance  int16  `db:"balance" json:"balance"`
}
