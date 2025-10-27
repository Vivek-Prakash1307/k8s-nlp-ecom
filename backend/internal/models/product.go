package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Rating      float64            `bson:"rating" json:"rating"`
	Popularity  int                `bson:"popularity" json:"popularity"`
	Category    string             `bson:"category" json:"category"`
	CreatedAt   int64              `bson:"createdAt" json:"createdAt"`
}
