package product

import "github.com/gin-gonic/gin"

func Route(router *gin.Engine, pc ProductController) {
	routes := router.Group("/api/v1/product")
	{
		routes.POST("", pc.CreateProduct)
		routes.GET("/:id", pc.GetProductDetail)
		routes.PUT("/:id", pc.UpdateProduct)
		routes.DELETE("/:id", pc.DeleteProduct)
	}
}
