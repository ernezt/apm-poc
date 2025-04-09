package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"apm/internal/api/handlers"
	"apm/internal/config"
	"apm/internal/db"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	config config.Config
	router *gin.Engine
	db     *db.Database
	server *http.Server
	logger *log.Logger

	// Services and handlers
	services    *services.Services
	handlers    *handlers.Factory
	authHandler *handlers.AuthHandler
}

// NewServer creates a new HTTP server
func NewServer(config config.Config, database *db.Database, logger *log.Logger) *Server {
	// Set Gin mode based on environment
	if config.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.Server.Environment == "test" {
		gin.SetMode(gin.TestMode)
	}

	s := &Server{
		config: config,
		db:     database,
		logger: logger,
	}

	// Initialize services
	s.initServices()

	// Initialize handlers
	s.initHandlers()

	// Set up router
	s.router = s.setupRoutes()

	// Create HTTP server
	s.server = &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      s.router,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,
	}

	return s
}

// initServices initializes all service instances
func (s *Server) initServices() {
	s.services = services.NewServices(s.db, s.logger)
}

// initHandlers initializes all handler instances
func (s *Server) initHandlers() {
	s.handlers = handlers.NewFactory(
		s.services.UserService,
		s.services.UserGroupService,
		s.services.StakeholderService,
		s.services.EntityService,
		s.services.SoftwareService,
		s.services.FunctionalCategoryService,
		s.services.SoftwareGroupService,
		s.services.StatusService,
		s.services.StatusLogService,
		s.services.RankService,
		s.services.NewsArticleService,
		s.services.MediaService,
		s.services.ProductDocumentationService,
		s.services.LogService,
	)

	// Initialize auth handler
	s.authHandler = handlers.NewAuthHandler(s.services.UserService, s.config.Server.JWTSecret)
}

// Start starts the HTTP server
func (s *Server) Start() error {
	s.logger.Printf("Starting server on port %s", s.config.Server.Port)
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}

// setupRoutes sets up the HTTP routes
func (s *Server) setupRoutes() *gin.Engine {
	// Create a gin router with default middleware
	router := gin.New()

	// Add custom middleware
	router.Use(s.loggerMiddleware())
	router.Use(s.corsMiddleware())
	router.Use(gin.Recovery()) // Recover from panics

	// Health check endpoint
	router.GET("/health", s.handleHealth)

	// API routes
	api := router.Group("/api")
	{
		// Register auth routes
		s.authHandler.RegisterRoutes(api)

		// Register all API routes using the handler factory
		s.handlers.RegisterRoutes(api)
	}

	return router
}

// handleHealth handles the health check route
func (s *Server) handleHealth(c *gin.Context) {
	if err := s.db.Ping(c.Request.Context()); err != nil {
		s.logger.Printf("Health check failed: %v", err)
		c.String(http.StatusServiceUnavailable, "Database connection error")
		return
	}
	c.String(http.StatusOK, "OK")
}

// loggerMiddleware creates a Gin middleware for logging requests
func (s *Server) loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		latency := time.Since(start)
		s.logger.Printf(
			"%s %s %s %s %d",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			latency,
			c.Writer.Status(),
		)
	}
}

// corsMiddleware adds CORS headers to responses
func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		origins := "*"
		if len(s.config.CORS.AllowedOrigins) > 0 && s.config.CORS.AllowedOrigins[0] != "*" {
			origins = s.config.CORS.AllowedOrigins[0]
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origins)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
