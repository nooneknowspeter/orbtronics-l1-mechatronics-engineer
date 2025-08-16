package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	DB *sql.DB
)

func SetupDatabaseConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", connectionString)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file:///migrations",
	// 	"postgres", driver)
	// m.Up()

	return db, nil
}

func StartDatabaseConnection() {
	connectionString := os.Getenv("POSTGRES_CONNECTION_URL")
	if connectionString == "" {
		log.Fatalf("database connection string not provided")
	}
	DB, _ = SetupDatabaseConnection(connectionString)
}
