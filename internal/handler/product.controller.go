package handler

import (
	"ecommerce_react_gin/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: service.NewProductService(),
	}
}

func (pc *ProductController) GetHomeProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "get all products",
	})
}
