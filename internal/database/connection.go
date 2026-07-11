package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/config"
)

func NewConnection(cfg *config.Config) (*sql.DB, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
