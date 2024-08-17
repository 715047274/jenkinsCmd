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

const schemaVersion = 1

func ensureSchema(db *sql.DB) {
	//sourceInstance, err := httpfs.New(http.FS(migrations), "migrations")
	//if err != nil {
	//	return fmt.Errorf("invalid source instance, %w", err)
	//}
	//targetInstance, err := sqlite.WithInstance(db, new(sqlite.Config))
	//if err != nil {
	//	return fmt.Errorf("invalid target sqlite instance, %w", err)
	//}
	//m, err := migrate.NewWithInstance(
	//	"httpfs", sourceInstance, "sqlite", targetInstance)
	//if err != nil {
	//	return fmt.Errorf("failed to initialize migrate instance, %w", err)
	//}
	//err = m.Migrate(schemaVersion)
	//if err != nil && err != migrate.ErrNoChange {
	//	return err
	//}
	//return sourceInstance.Close()
	//////////------- old---
	//migration, err := migrate.New("db/migration", "sqlite3://sqliteDemo.db")
	//if err != nil {
	//	// log.Fatal().Err(err).Msg("cannot create new migrate instance")
	//	fmt.Errorf("invalid source instance, %w", err)
	//}
	//
	//if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
	//	// log.Fatal().Err(err).Msg("failed to run migrate up")
	//	fmt.Errorf("failed to initialize migrate instance, %w", err)
	//}

	//log.Info().Msg("db migrated successfully")
	//////////------- old--- end
	// Get database driver for SQLite

	// Create a new migration instance

}
func init() {

	//db, err := sql.Open("sqlite3", "./sqliteDemo.db")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ensureSchema(db)
	//fmt.Println("DB ..... Created")
	////if err := ensureSchema(db); err != nil {
	////	log.Fatalln("migration failed")
	////}
	//DB = db
}

//func setupRoute() *gin.Engine {
//	r := gin.Default()
//	r.GET("/getAccounts", getAccounts)
//	return r
//}

//	func getAccounts(c *gin.Context) {
//		var sql = `SELECT * FROM accounts`
//		rows, _ := DB.Query(sql)
//		defer rows.Close()
//
//		accounts := make([]Account, 0)
//		for rows.Next() {
//			singleAccount := Account{}
//			_ = rows.Scan(&singleAccount.Id, &singleAccount.Balance, &singleAccount.Currency, &singleAccount.Owner)
//
//			accounts = append(accounts, singleAccount)
//		}
//		_ = rows.Err()
//
//		fmt.Println(rows)
//
//		c.JSON(http.StatusOK, accounts)
//	}
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
	//serv := setupRoute()
	//serv.Run(":8080")
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
