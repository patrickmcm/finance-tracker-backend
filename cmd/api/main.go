package main

import (
	"database/sql"
	v1 "finance-tracker-backend/internal/api/v1"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	DBUser string
	DBPass string
}

func LoadConfig() *Config {
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER ENV VAR REQUIRED")
	}
	password := os.Getenv("DB_PASS")
	if password == "" {
		log.Fatal("DB_USER ENV VAR REQUIRED")
	}

	return &Config{
		DBUser: user,
		DBPass: password,
	}
}

func main() {
	cfg := LoadConfig()

	app := NewApp(cfg)
	log.Fatal(app.Start())
}

type App struct {
	cfg *Config
}

func NewApp(cfg *Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() error {
	dbCfg := pq.Config{
		Host:     "localhost",
		Port:     5432,
		SSLMode:  pq.SSLModeDisable,
		Database: "finance_tracker",
		User:     a.cfg.DBUser,
		Password: a.cfg.DBPass,
	}

	connCfg, err := pq.NewConnectorConfig(dbCfg)
	if err != nil {
		return err
	}

	db := sql.OpenDB(connCfg)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	v1API := v1.NewRoute(db)

	v1API.RegisterHandlers(mux)

	srv := http.Server{
		Addr:         ":3000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server running on port %s\n", srv.Addr)
	err = srv.ListenAndServe()
	return err
}
