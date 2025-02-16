package controller

import (
	"ecommerce_react_gin/internal/models"
	"ecommerce_react_gin/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		productService: service.NewProductService(),
	}
}

// [GET] /v1/api/product
func (pc *ProductController) GetAllProducts(c *gin.Context) {
	products, err := pc.productService.FindAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

// [POST] /v1/api/product/add
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.ProductModel
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product.ID = primitive.NewObjectID()

	if err := pc.productService.CreateOneProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"msg": "product created successfully"})
}

// [GET] /v1/api/product/:id
func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.productService.FindProductById(objectID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, gin.H{"product": product})
}

// [PATCH] /v1/api/product/:id
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.ProductModel
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product.ID = objectID

	if err := pc.productService.UpdateOneById(objectID, &product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "product updated successfully"})
}

// [DELETE] /v1/api/product/:id
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := pc.productService.DeleteOneById(objectID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "product deleted successfully"})
}
