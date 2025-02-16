package service

import (
	"ecommerce_react_gin/internal/models"
	"ecommerce_react_gin/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductService interface {
	FindProductById(id interface{}) (*models.ProductModel, error)
	FindAllProducts() ([]models.ProductModel, error)
	CreateOneProduct(product *models.ProductModel) error
	UpdateOneById(id interface{}, update interface{}) error
	DeleteOneById(id interface{}) error
}

type productService struct {
	productRepo repo.ProductRepo
}

func NewProductService() ProductService {
	return &productService{
		productRepo: repo.NewProductRepo(),
	}
}

func (ps *productService) FindProductById(id interface{}) (*models.ProductModel, error) {
	filter := bson.M{"_id": id}
	return ps.productRepo.FindOne(filter)
}

func (ps *productService) FindAllProducts() ([]models.ProductModel, error) {
	filter := bson.M{}
	return ps.productRepo.FindMany(filter)
}

func (ps *productService) CreateOneProduct(product *models.ProductModel) error {
	return ps.productRepo.Create(product)
}

func (ps *productService) UpdateOneById(id interface{}, update interface{}) error {
	filter := bson.M{"_id": id}
	updateDoc := bson.M{"$set": update}
	return ps.productRepo.Update(filter, updateDoc)
}

func (ps *productService) DeleteOneById(id interface{}) error {
	filter := bson.M{"_id": id}
	return ps.productRepo.Delete(filter)
}
