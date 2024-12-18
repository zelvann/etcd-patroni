package product

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/zelvann/minio-ds/internal/domain/image"
	"github.com/zelvann/minio-ds/internal/utils"
)

const (
	bucketName = "images"
)

type Controller struct {
	repo        Repository
	minioClient *minio.Client
	imgCtrl     *image.Controller
}

func NewController(repo Repository, minioClient *minio.Client) *Controller {
	return &Controller{
		repo:        repo,
		minioClient: minioClient,
		imgCtrl:     image.NewController(minioClient),
	}
}

func (c *Controller) CreateProduct(ctx *gin.Context) {
	// Get form data first
	name := ctx.PostForm("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Name is required", nil))
		return
	}

	description := ctx.PostForm("description")

	// Parse expiry date if provided
	var expiryDate *time.Time
	if expiry := ctx.PostForm("expiry_date"); expiry != "" {
		parsedTime, err := time.Parse(time.RFC3339, expiry)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Invalid expiry date format", err.Error()))
			return
		}
		expiryDate = &parsedTime
	}

	// Handle image upload
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Failed to get image", err.Error()))
		return
	}
	defer file.Close()

	// Generate unique filename
	filename := filepath.Base(header.Filename)

	// Upload to MinIO using image controller's logic
	_, err = c.minioClient.PutObject(ctx, image.BucketName, filename, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})
	if err != nil {
		fmt.Printf("Failed to upload to MinIO: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to upload image", err.Error()))
		return
	}

	// Create the MinIO URL for the image (using the bucket name from image package)
	imagePath := fmt.Sprintf("20.198.248.111:9000/%s/%s", image.BucketName, filename)

	// Create product with image path
	product := &Product{
		Name:        name,
		Description: description,
		ImagePath:   imagePath,
		ExpiryDate:  expiryDate,
	}

	if err := c.repo.Create(ctx, product); err != nil {
		// If product creation fails, try to delete the uploaded image
		_ = c.minioClient.RemoveObject(ctx, image.BucketName, filename, minio.RemoveObjectOptions{})
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to create product", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewSucessResponse("Product created successfully").WithPayload(product))
}

func (c *Controller) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Invalid product ID", err.Error()))
		return
	}

	product, err := c.repo.FindByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewFailedResponse("Product not found", err.Error()))
		return
	}

	var req UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Invalid request body", err.Error()))
		return
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.ExpiryDate != nil {
		product.ExpiryDate = req.ExpiryDate
	}

	if err := c.repo.Update(ctx, product); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to update product", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Product updated successfully").WithPayload(product))
}

func (c *Controller) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Invalid product ID", err.Error()))
		return
	}

	product, err := c.repo.FindByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewFailedResponse("Product not found", err.Error()))
		return
	}

	// Extract filename from image path
	filename := filepath.Base(product.ImagePath)

	// Delete image from MinIO
	err = c.minioClient.RemoveObject(ctx, image.BucketName, filename, minio.RemoveObjectOptions{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to delete image", err.Error()))
		return
	}

	if err := c.repo.Delete(ctx, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to delete product", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Product deleted successfully"))
}

func (c *Controller) GetProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewFailedResponse("Invalid product ID", err.Error()))
		return
	}

	product, err := c.repo.FindByID(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewFailedResponse("Product not found", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Product retrieved successfully").WithPayload(product))
}

func (c *Controller) ListProducts(ctx *gin.Context) {
	products, err := c.repo.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewFailedResponse("Failed to retrieve products", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSucessResponse("Products retrieved successfully").WithPayload(products))
}
