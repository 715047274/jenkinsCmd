package db

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "log"
)

func MigrateDatabase(dbPath string) error {
	// Get database driver for SQLite
	//driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	//if err != nil {
	//	log.Fatalf("Could not create SQLite driver: %v", err)
	//}

	// Create a new migration instance
	//m, err := migrate.NewWithDatabaseInstance("file://migration",
	//	"sqlite3", driver)
	//if err != nil {
	//	log.Fatalf("Migration initialization failed: %v", err)
	//}
	////
	////// Run migration UP
	//if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	//	log.Fatalf("Could not apply migrations: %v", err)
	//} else if err == migrate.ErrNoChange {
	//	log.Println("No new migrations to apply")
	//} else {
	//	log.Println("Migrations applied successfully")
	//}
	return nil
}
