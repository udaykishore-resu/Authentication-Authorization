package database

import (
	"database/sql"
	"session-cookie/config"

	_ "github.com/lib/pq"
)

func NewConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DatabaseConnectionString())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
