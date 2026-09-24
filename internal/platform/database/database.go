package database

import (
	"context"
	"finance-tracker-backend/internal/platform/config"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s dbname=finance_tracker sslmode=disable", cfg.DBUser, cfg.DBPass))
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
