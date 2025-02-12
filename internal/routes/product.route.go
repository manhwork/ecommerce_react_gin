package routes

import (
	"ecommerce_react_gin/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(v1 *gin.RouterGroup) {
	ProductRouter := v1.Group("/product")
	{
		ProductRouter.GET("/", controller.NewProductController().GetHomeProduct)
	}
}
