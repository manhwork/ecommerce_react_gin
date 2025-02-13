package routes

import (
	"ecommerce_react_gin/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(v1 *gin.RouterGroup) {
	ProductRouter := v1.Group("/product")
	{
		ProductRouter.GET("/", controller.NewProductController().GetAllProducts)
		ProductRouter.GET("/:id", controller.NewProductController().GetProduct)
		ProductRouter.POST("/add", controller.NewProductController().CreateProduct)
		ProductRouter.PATCH("/:id", controller.NewProductController().UpdateProduct)
		ProductRouter.DELETE("/:id", controller.NewProductController().DeleteProduct)
	}
}
