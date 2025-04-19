package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/axz/control-plane/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server represents the API server
type Server struct {
	router *gin.Engine
	server *http.Server
	db     *gorm.DB
}

// NewServer creates a new API server
func NewServer(db *gorm.DB, cfg *config.ServerConfig) *Server {
	router := gin.Default()
	
	server := &Server{
		router: router,
		db:     db,
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      router,
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		},
	}

	server.setupRoutes()
	return server
}

// setupRoutes configures the API routes
func (s *Server) setupRoutes() {
	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 group
	v1 := s.router.Group("/api/v1")
	{
		serviceHandler := NewServiceHandler(s.db)
		v1.POST("/services", serviceHandler.RegisterService)
	}
}

// Start starts the API server
func (s *Server) Start() error {
	log.Printf("Starting API server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down API server...")
	return s.server.Shutdown(ctx)
} 