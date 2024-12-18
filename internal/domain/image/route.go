package image

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func Route(router *gin.Engine, minioClient *minio.Client) {
	controller := NewController(minioClient)

	group := router.Group("/images")
	{
		group.POST("/upload", controller.UploadImage)
		group.GET("/list", controller.ListImages)
		group.GET("/:filename", controller.GetImage)
		group.DELETE("/:filename", controller.DeleteImage)
	}
}
