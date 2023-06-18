package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/dannylee/url-ports-adapters/adapters"
	"github.com/dannylee/url-ports-adapters/application/api"

	_ "github.com/lib/pq"
)

func main() {
	config := adapters.Config{
		Port: 8080,
		Env:  "dev",
		DB: struct {
			Dsn            string
			MaxIdleTimeout string
		}{
			Dsn:            os.Getenv("DB_DSN"),
			MaxIdleTimeout: "5m",
		},
		Limiter: false,
	}

	db, err := openDB(config)
	if err != nil {
		log.Printf("error initializing database: %v", err)
		return
	}

	urlAdapter := adapters.New(config, db)
	urlController := api.NewUrlController(urlAdapter)

	err = serve(config, api.Router(urlController))
	if err != nil {
		log.Printf("Error starting server at :%v", config.Port)
	}
}

func openDB(cfg adapters.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(cfg.DB.MaxIdleTimeout)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
