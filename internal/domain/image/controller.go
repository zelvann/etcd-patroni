package image

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/zelvann/minio-ds/internal/utils"
)

const (
	bucketName = "images"
)

type Controller struct {
	minioClient *minio.Client
}

func NewController(minioClient *minio.Client) *Controller {
	return &Controller{
		minioClient: minioClient,
	}
}

// UploadImage handles image upload
func (c *Controller) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Failed to get image", err.Error()))
		return
	}
	defer file.Close()

	filename := filepath.Base(header.Filename)

	_, err = c.minioClient.PutObject(ctx, bucketName, filename, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to upload image", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Image uploaded successfully"))
}

func (c *Controller) GetImage(ctx *gin.Context) {
	filename := ctx.Param("filename")

	object, err := c.minioClient.GetObject(ctx, bucketName, filename, minio.GetObjectOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to get image", err.Error()))
		return
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.DataFromReader(http.StatusOK, -1, "application/octet-stream", object, nil)
}

func (c *Controller) ListImages(ctx *gin.Context) {
	var images []string

	objectCh := c.minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to list images", object.Err.Error()))
			return
		}
		images = append(images, object.Key)
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Images retrieved successfully").WithPayload(images))
}

func (c *Controller) DeleteImage(ctx *gin.Context) {
	filename := ctx.Param("filename")

	err := c.minioClient.RemoveObject(ctx, bucketName, filename, minio.RemoveObjectOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to delete image", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Image deleted successfully"))
}
