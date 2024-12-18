package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/zelvann/minio-ds/internal/config"
	"github.com/zelvann/minio-ds/internal/domain/icmp"
	"github.com/zelvann/minio-ds/internal/domain/image"
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

	err = minioClient.MakeBucket(context.Background(), "images", minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), "images")
		if errBucketExists == nil && exists {
			log.Printf("Bucket 'images' already exists")
		} else {
			log.Fatal(err)
		}
	}

	image.Route(server, minioClient)

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
