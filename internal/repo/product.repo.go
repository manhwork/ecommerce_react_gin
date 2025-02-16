package repo

import (
	"context"
	"ecommerce_react_gin/internal/database"
	"ecommerce_react_gin/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo interface {
	FindMany(filter interface{}) ([]models.ProductModel, error)
	FindOne(filter interface{}) (*models.ProductModel, error)
	Create(product *models.ProductModel) error
	Update(filter interface{}, update interface{}) error
	Delete(filter interface{}) error
}

type productRepo struct {
	ProductCollection *mongo.Collection
}

func NewProductRepo() *productRepo {
	return &productRepo{
		ProductCollection: database.New().GetCollection("product"),
	}
}

func (pr *productRepo) FindMany(filter interface{}) ([]models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productsList := []models.ProductModel{}

	cursor, err := pr.ProductCollection.Find(ctx, filter)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for cursor.Next(ctx) {
		var product models.ProductModel

		err = cursor.Decode(&product)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		productsList = append(productsList, product)
	}

	return productsList, nil
}

func (pr *productRepo) FindOne(filter interface{}) (*models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var product models.ProductModel

	err := pr.ProductCollection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil
}

func (pr *productRepo) Create(product *models.ProductModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := pr.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepo) Update(filter interface{}, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := pr.ProductCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pr *productRepo) Delete(filter interface{}) error {
	_, err := pr.ProductCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
