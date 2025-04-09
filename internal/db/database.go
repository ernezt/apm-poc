package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Database represents the database connection pool
type Database struct {
	Pool *pgxpool.Pool
}

// Config holds database configuration options
type Config struct {
	URL            string
	MaxConns       int32
	MinConns       int32
	MaxConnLife    time.Duration
	MaxConnIdle    time.Duration
	ConnectTimeout time.Duration
}

// DefaultConfig returns a default database configuration
func DefaultConfig(url string) Config {
	return Config{
		URL:            url,
		MaxConns:       10,
		MinConns:       2,
		MaxConnLife:    time.Hour,
		MaxConnIdle:    5 * time.Minute,
		ConnectTimeout: 10 * time.Second,
	}
}

// New creates a new Database instance
func New(config Config) (*Database, error) {
	// Parse connection string
	poolConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing database URL: %w", err)
	}

	// Apply configuration
	poolConfig.MaxConns = config.MaxConns
	poolConfig.MinConns = config.MinConns
	poolConfig.MaxConnLifetime = config.MaxConnLife
	poolConfig.MaxConnIdleTime = config.MaxConnIdle

	// Create a timeout context for connection
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectTimeout)
	defer cancel()

	// Connect to database
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{Pool: pool}, nil
}

// Close closes the database connection pool
func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// Ping checks database connectivity
func (db *Database) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
