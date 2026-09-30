package controllers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

type Controller interface {
	RegisterController(s *grpc.Server, db *pgxpool.Pool)
}
