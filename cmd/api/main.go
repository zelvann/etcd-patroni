package main

import (
	"fmt"
	"log"

	"github.com/zelvann/minio-ds/internal/config"
	"github.com/zelvann/minio-ds/internal/domain/icmp"
	"github.com/zelvann/minio-ds/internal/instance"
	"github.com/zelvann/minio-ds/internal/middleware"
)

func main() {
	env := config.LoadEnv()

	minioClient, err := instance.NewMinio(env.MinioEndpoint, env.MinioAccessKey, env.MinioSecretKey)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}
	fmt.Println(minioClient)

	server := instance.NewGin()
	server.Use(middleware.CORS())

	icmp.Route(server)

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
