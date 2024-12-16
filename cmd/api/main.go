package main

import (
	"log"

	"github.com/zelvann/minio-ds/internal/config"
	"github.com/zelvann/minio-ds/internal/domain/ping"
	"github.com/zelvann/minio-ds/internal/instance"
	"github.com/zelvann/minio-ds/internal/middleware"
)

func main() {
	env := config.LoadEnv()

	server := instance.NewGin()
	server.Use(middleware.CORS())

	ping.Route(server) // icmp

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
