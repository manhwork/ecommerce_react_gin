package service

import (
	"ecommerce_react_gin/internal/models"
	"ecommerce_react_gin/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repo.NewProductRepo(),
	}
}

func (ps *ProductService) FindProductById(id interface{}) (*models.ProductModel, error) {
	filter := bson.M{"_id": id}
	return ps.productRepo.FindOne(filter)
}

func (ps *ProductService) FindAllProducts() ([]models.ProductModel, error) {
	filter := bson.M{}
	return ps.productRepo.FindMany(filter)
}

func (ps *ProductService) CreateOneProduct(product *models.ProductModel) error {
	return ps.productRepo.Create(product)
}

func (ps *ProductService) UpdateOneById(id interface{}, update interface{}) error {
	filter := bson.M{"_id": id}
	updateDoc := bson.M{"$set": update}
	return ps.productRepo.Update(filter, updateDoc)
}

func (ps *ProductService) DeleteOneById(id interface{}) error {
	filter := bson.M{"_id": id}
	return ps.productRepo.Delete(filter)
}
