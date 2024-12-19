package main

import (
	"log"

	"github.com/zelvann/minio-ds/internal/config"
	"github.com/zelvann/minio-ds/internal/domain/icmp"
	"github.com/zelvann/minio-ds/internal/domain/product"
	"github.com/zelvann/minio-ds/internal/instance"
	"github.com/zelvann/minio-ds/internal/middleware"
)

func main() {
	env := config.LoadEnv()

	db := instance.NewPostgres(env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)
	if err := db.AutoMigrate(&product.ProductSchema{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	minioClient, err := instance.NewMinio(env.MinioEndpoint, env.MinioAccessKey, env.MinioSecretKey)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	productRepository := product.NewProductRepository(db)
	productUsecase := product.NewProductUsecase(productRepository, minioClient)
	productController := product.NewProductController(productUsecase)

	server := instance.NewGin()
	server.Use(middleware.CORS())

	product.Route(server, productController)
	icmp.Route(server)

	if env.ApiPort == "" {
		env.ApiPort = "8080"
	}

	if err := server.Run(":" + env.ApiPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
