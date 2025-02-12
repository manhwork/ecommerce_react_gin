package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductModel struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
	Price string             `bson:"price" json:"price"`
}
