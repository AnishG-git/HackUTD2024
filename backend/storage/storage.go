package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	username      = "postgres"
	password      = "postgres"
	pgcontainer   = "db"
	migrationsDir = "./migrations"
	port          = 5432
)

func LoadDB(dbname string, logger *log.Logger) (*sql.DB, error) {
	// Load the database
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", username, password, dbname, pgcontainer, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
