package main

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	repo "github.com/JeffreyOmoakah/E-commerce-backend-API/internal/adapters/postgresql/sqlc"
	"github.com/JeffreyOmoakah/E-commerce-backend-API/internal/orders"
	"github.com/JeffreyOmoakah/E-commerce-backend-API/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Mount routes to the application
func (app *application) mount () http.Handler {
 r := chi.NewRouter()
 // A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes
 
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	
	// The health endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})
 	
	productRepo := repo.New(app.db)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	orderRepo := repo.New(app.db)
	orderService := orders.NewService(orderRepo,app.db, app.logger)
	ordersHandler := orders.NewHandler(orderService)


	// Create v1 route group
	r.Route("/v1", func(r chi.Router) {
		r.Get("/products", productHandler.ListProducts)
		r.Post("/orders", ordersHandler.PlaceOrder)
	})

	return r
}

// Run the application 
func (app *application) run(h http.Handler) error{
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: time.Minute,
	}
	
	// print a message to the console
	log.Printf("Starting server on %s", app.config.addr)
	
	return srv.ListenAndServe()
}

type application struct {
	config config // Application manager
	db *pgxpool.Pool
	logger *slog.Logger
}

type config struct {
	addr string // Address of the API server 
	db dbConfig // Database configuration
}

type dbConfig struct {
	dsn string // Database connection string 
}