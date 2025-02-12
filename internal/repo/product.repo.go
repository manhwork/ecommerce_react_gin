package repo

import (
	"context"
	"ecommerce_react_gin/internal/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	ProductCollection *mongo.Collection
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		ProductCollection: database.New().GetCollection("product"),
	}
}

func (pr *ProductRepo) FindMany() ([]interface{}, error) {
	cursor, err := pr.ProductCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []interface{}
	if err = cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepo) Create(product interface{}) (*mongo.InsertOneResult, error) {
	result, err := pr.ProductCollection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *ProductRepo) FindOne(filter interface{}) (interface{}, error) {
	var product interface{}
	err := pr.ProductCollection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductRepo) Update(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	result, err := pr.ProductCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *ProductRepo) Delete(filter interface{}) (*mongo.DeleteResult, error) {
	result, err := pr.ProductCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
