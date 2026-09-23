package database

import (
	"database/sql"
	"finance-tracker-backend/internal/platform/config"
	"github.com/lib/pq"
)

func New(cfg *config.Config) (*sql.DB, error) {
	dbCfg := pq.Config{
		Host:     "localhost",
		Port:     5432,
		SSLMode:  pq.SSLModeDisable,
		Database: "finance_tracker",
		User:     cfg.DBUser,
		Password: cfg.DBPass,
	}

	connCfg, err := pq.NewConnectorConfig(dbCfg)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connCfg)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
