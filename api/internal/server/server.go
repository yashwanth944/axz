package server

import (
    "github.com/labstack/echo/v4"
    "axz/api/internal/config"
    "axz/api/internal/handlers"
    "axz/api/internal/store"
)

type Server struct {
    echo   *echo.Echo
    config *config.Config
    store  *store.Store
}

func New(cfg *config.Config) (*Server, error) {
    e := echo.New()
    
    // Initialize store
    store, err := store.New(cfg.Database)
    if err != nil {
        return nil, err
    }

    s := &Server{
        echo:   e,
        config: cfg,
        store:  store,
    }

    // Register routes
    s.registerRoutes()

    return s, nil
}

func (s *Server) Start() error {
    return s.echo.Start(s.config.ListenAddr)
}

func (s *Server) registerRoutes() {
    // Health check
    s.echo.GET("/health", handlers.Health)

    // API v1
    v1 := s.echo.Group("/api/v1")
    
    // Services
    v1.POST("/services", handlers.CreateService(s.store))
    v1.GET("/services", handlers.ListServices(s.store))
    v1.GET("/services/:id", handlers.GetService(s.store))
    v1.PUT("/services/:id", handlers.UpdateService(s.store))
    v1.DELETE("/services/:id", handlers.DeleteService(s.store))

    // Routes
    v1.POST("/routes", handlers.CreateRoute(s.store))
    v1.GET("/routes", handlers.ListRoutes(s.store))
    v1.GET("/routes/:id", handlers.GetRoute(s.store))
    v1.PUT("/routes/:id", handlers.UpdateRoute(s.store))
    v1.DELETE("/routes/:id", handlers.DeleteRoute(s.store))
} 