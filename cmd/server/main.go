package main

import (
	instrumentsController "finance-tracker-backend/internal/controllers/instruments"
	"finance-tracker-backend/internal/platform/config"
	"finance-tracker-backend/internal/platform/database"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := config.LoadConfig()

	app := NewApp(cfg)
	log.Fatal(app.Start())
}

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() error {
	db, err := database.New(a.cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	reflection.Register(srv)

	instruments := instrumentsController.New(db)
	instruments.RegisterController(srv)
	err = srv.Serve(lis)
	return err
}
