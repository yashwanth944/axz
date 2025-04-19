package main

import (
    "log"
    "github.com/axz/api/internal/server"
    "github.com/axz/api/internal/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    srv, err := server.New(cfg)
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }

    if err := srv.Start(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
} 