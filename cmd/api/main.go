package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
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

	// Check MinIO connection
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("Failed to list buckets: %v", err)
	}
	fmt.Printf("Connected to MinIO. Available buckets: %v\n", buckets)

	db := instance.NewPostgres(env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)

	// Auto-migrate the product model
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	server := instance.NewGin()
	server.Use(middleware.CORS())

	icmp.Route(server)

	// Create images bucket if it doesn't exist
	err = minioClient.MakeBucket(context.Background(), "images", minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), "images")
		if errBucketExists == nil && exists {
			log.Printf("Bucket 'images' already exists")
		} else {
			log.Fatalf("Failed to create bucket: %v", err)
		}
	} else {
		log.Printf("Successfully created 'images' bucket")
	}

	image.Route(server, minioClient)
	product.Route(server, db, minioClient)

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
