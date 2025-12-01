package main

import (
	"log/slog"
	"os"

	"github.com/JeffreyOmoakah/E-commerce-backend-API/internal/env"
)

func main() {
	cfg:= config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{},
	}
	api := application{
		config: cfg,
	}
	
	api.run(api.mount())
	
	// Logger
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	
	
		if err := api.run(api.mount()); err != nil {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
}