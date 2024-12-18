package main

import (
	"log"

	"github.com/zelvann/minio-ds/internal/config"
	"github.com/zelvann/minio-ds/internal/domain/icmp"
	"github.com/zelvann/minio-ds/internal/domain/image"
	"github.com/zelvann/minio-ds/internal/domain/product"
	"github.com/zelvann/minio-ds/internal/instance"
	"github.com/zelvann/minio-ds/internal/middleware"
)

func main() {
	env := config.LoadEnv()

	minioClient, err := instance.NewMinio(env.MinioEndpoint, env.MinioAccessKey, env.MinioSecretKey)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	db := instance.NewPostgres(env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	server := instance.NewGin()
	server.Use(middleware.CORS())

	icmp.Route(server)
	image.Route(server, minioClient)
	product.Route(server, db, minioClient)

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
