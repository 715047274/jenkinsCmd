package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var DB *sql.DB
var migrations embed.FS

const schemaVersion = 1

func ensureSchema() {
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
	migration, err := migrate.New("db/migration", "sqlite3://sqliteDemo.db")
	if err != nil {
		// log.Fatal().Err(err).Msg("cannot create new migrate instance")
		fmt.Errorf("invalid source instance, %w", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		// log.Fatal().Err(err).Msg("failed to run migrate up")
		fmt.Errorf("failed to initialize migrate instance, %w", err)
	}

	//log.Info().Msg("db migrated successfully")
}
func init() {
	ensureSchema()
	db, err := sql.Open("sqlite3", "./sqliteDemo.db")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DB ..... Created")
	//if err := ensureSchema(db); err != nil {
	//	log.Fatalln("migration failed")
	//}
	DB = db
}

func setupRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/getAccounts", getAccounts)
	return r
}

func getAccounts(c *gin.Context) {
	var sql = `SELECT * FROM accounts`
	rows, _ := DB.Query(sql)
	defer rows.Close()

	accounts := make([]Account, 0)
	for rows.Next() {
		singleAccount := Account{}
		_ = rows.Scan(&singleAccount.Id, &singleAccount.Balance, &singleAccount.Currency, &singleAccount.Owner)

		accounts = append(accounts, singleAccount)
	}
	_ = rows.Err()

	fmt.Println(rows)

	c.JSON(http.StatusOK, accounts)
}

func main() {
	serv := setupRoute()
	serv.Run(":8080")
	fmt.Println("hello world")
}

type Account struct {
	Id       int    `db:"id" json:"id"`
	Owner    string `db:"owner" json:"owner"`
	Currency string `db:"currency" json:"currency"`
	Balance  int16  `db:"balance" json:"balance"`
}
