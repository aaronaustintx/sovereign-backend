package database

import (
	"database/sql"

	"github.com/aaronaustintx/sovereign-backend/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
