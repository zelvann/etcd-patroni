package product

import (
	"errors"
	"mime/multipart"
	"time"
)

const (
	PRODUCT_BUCKET = "product"

	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed to get data from body"
	MESSAGE_FAILED_CREATE_PRODUCT     = "failed to create product"
	MESSAGE_FAILED_GET_PRODUCT        = "failed to get product"
	MESSAGE_FAILED_UPDATE_PRODUCT     = "failed to update product"
	MESSAGE_FAILED_DELETE_PRODUCT     = "failed to delete product"

	MESSAGE_SUCCESS_CREATE_PRODUCT = "success create product"
	MESSAGE_SUCCESS_GET_PRODUCT    = "success get product"
	MESSAGE_SUCCESS_UPDATE_PRODUCT = "success update product"
	MESSAGE_SUCCESS_DELETE_PRODUCT = "success delete product"
)

var (
	ErrFormatFileNotValid      = errors.New("format file not valid")
	ErrProductNotFound         = errors.New("product not found")
	ErrSomethingWentWrong      = errors.New("something went wrong")
	ErrMinioSomethingWentWrong = errors.New("something went wrong in minio")
	ErrBucketNotFound          = errors.New("bucket not found")
)

type (
	CreateProductDTO struct {
		ID             uint                  `json:"id" form:"id" binding:"required"`
		Name           string                `json:"name" form:"name" binding:"required"`
		Description    string                `json:"description" form:"description" binding:"required"`
		ProductPicture *multipart.FileHeader `json:"product_picture" form:"product_picture" binding:"required"`
		ExpireDate     time.Time             `json:"expire_date" form:"expire_date" binding:"required"`
	}

	GetProductDTO struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ExpireDate  string `json:"expire_date"`
	}

	UpdateProductDTO struct {
		Name        string `json:"name" form:"name" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
	}
)
