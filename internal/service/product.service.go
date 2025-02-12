package service

import "ecommerce_react_gin/internal/repo"

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repo.NewProductRepo(),
	}
}
