package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zelvann/minio-ds/internal/utils"
)

type (
	ProductController interface {
		CreateProduct(*gin.Context)
		GetProductDetail(*gin.Context)
		UpdateProduct(*gin.Context)
		DeleteProduct(*gin.Context)
	}

	productController struct {
		uc ProductUsecase
	}
)

func NewProductController(uc ProductUsecase) ProductController {
	return &productController{uc: uc}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var req CreateProductDTO
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.uc.CreateProduct(ctx.Request.Context(), req); err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_CREATE_PRODUCT, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewSucessResponse(MESSAGE_SUCCESS_CREATE_PRODUCT)
	ctx.JSON(http.StatusCreated, res)
}

func (c *productController) GetProductDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.uc.GetProductByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_GET_PRODUCT, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewSucessResponse(MESSAGE_SUCCESS_GET_PRODUCT).WithPayload(result)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var req UpdateProductDTO
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.uc.UpdateProduct(ctx.Request.Context(), req, id); err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_UPDATE_PRODUCT, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewSucessResponse(MESSAGE_SUCCESS_UPDATE_PRODUCT)
	ctx.JSON(http.StatusOK, res)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.uc.DeleteProduct(ctx.Request.Context(), id); err != nil {
		res := utils.NewFailedResponse(MESSAGE_FAILED_DELETE_PRODUCT, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.NewSucessResponse(MESSAGE_SUCCESS_DELETE_PRODUCT)
	ctx.JSON(http.StatusOK, res)
}
