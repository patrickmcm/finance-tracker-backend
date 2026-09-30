package main

import (
	"finance-tracker-backend/internal/controllers/instrumentprices"
	instrumentsController "finance-tracker-backend/internal/controllers/instruments"
	"finance-tracker-backend/internal/platform/config"
	"finance-tracker-backend/internal/platform/controllers"
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

	ctrls := []controllers.Controller{&instrumentsController.Server{}, &instrumentprices.Server{}}

	app := NewApp(cfg, ctrls)
	log.Fatal(app.Start())
}

type App struct {
	cfg   *config.Config
	ctrls []controllers.Controller
}

func NewApp(cfg *config.Config, ctrls []controllers.Controller) *App {
	return &App{cfg: cfg, ctrls: ctrls}
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

	for _, v := range a.ctrls {
		v.RegisterController(srv, db)
	}

	err = srv.Serve(lis)
	return err
}
