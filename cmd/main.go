package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/layardaputra/govtech-catalog-test-project/common"
	"github.com/layardaputra/govtech-catalog-test-project/config"
	"github.com/layardaputra/govtech-catalog-test-project/database"
	"github.com/layardaputra/govtech-catalog-test-project/internal/app"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize the database
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Failed to initialize the database: %v\n", err)
		return
	}

	// Initialize the application
	apps := app.NewApp(db)

	// Initialize the Chi router
	router := chi.NewRouter()

	// Middleware
	router.Use(common.ExceptionHandlerMiddleware)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes
	app.RegisterRoutes(apps, router)

	// Start the HTTP server
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: router,
	}

	// Set up a signal channel to capture interrupt signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server is running on port %s\n", serverAddr)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	// Wait for an interrupt signal
	<-signalCh

	// Perform graceful shutdown
	log.Println("Shutting down the server...")

	// Close the HTTP server using defer
	defer func() {
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server shutdown error: %v\n", err)
		}
	}()

	// Close the database connection using defer
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Database connection close error: %v\n", err)
		}
	}()

	log.Println("Server has been gracefully shut down.")
}
