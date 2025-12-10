package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/JeffreyOmoakah/E-commerce-backend-API/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	
	cfg:= config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn:env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}
	
	// Logger
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	
	
	// database 
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	
	logger.Info("connected to database","dsn",cfg.db.dsn)
	
	api := application{
		config: cfg,
	}
	
	api.run(api.mount())
	
		if err := api.run(api.mount()); err != nil {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
}