package product

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

func Route(router *gin.Engine, db *gorm.DB, minioClient *minio.Client) {
	repo := NewRepository(db)
	controller := NewController(repo, minioClient)

	group := router.Group("/products")
	{
		group.POST("", controller.CreateProduct)
		group.PUT("/:id", controller.UpdateProduct)
		group.DELETE("/:id", controller.DeleteProduct)
		group.GET("/:id", controller.GetProduct)
		group.GET("", controller.ListProducts)
	}
}
