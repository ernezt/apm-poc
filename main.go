package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apm/internal/api"
	"apm/internal/config"
	"apm/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// It's okay if .env doesn't exist, we'll use environment variables
		// or default values
	}

	// Parse command line flags
	showEnvHelp := flag.Bool("env-help", false, "Show environment variable configuration help")
	flag.Parse()

	// If --env-help flag is set, print usage and exit
	if *showEnvHelp {
		config.PrintUsage()
		return
	}

	// Set up logger
	logger := log.New(os.Stdout, "[APM] ", log.LstdFlags)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}
	logger.Printf("Starting Application Portfolio Management (APM) in %s mode", cfg.Server.Environment)

	// Connect to the database
	dbConfig := db.Config{
		URL:            cfg.Database.URL,
		MaxConns:       int32(cfg.Database.MaxConns),
		MinConns:       int32(cfg.Database.MinConns),
		MaxConnLife:    time.Duration(cfg.Database.MaxConnLife) * time.Second,
		MaxConnIdle:    time.Duration(cfg.Database.MaxConnIdle) * time.Second,
		ConnectTimeout: time.Duration(cfg.Database.ConnectTimeout) * time.Second,
	}

	database, err := db.New(dbConfig)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Test database connection
	if err := database.Ping(context.Background()); err != nil {
		logger.Fatalf("Failed to ping database: %v", err)
	}
	logger.Println("Connected to database successfully")

	// Initialize the server
	server := api.NewServer(cfg, database, logger)

	// Start the server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Create a deadline context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Println("Server exited gracefully")
}
