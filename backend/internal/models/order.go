package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Items     []CartItem         `bson:"items" json:"items"`
	Total     float64            `bson:"total" json:"total"`
	Status    string             `bson:"status" json:"status"` // pending, shipped, delivered
	Address   string             `bson:"address" json:"address"`
	CreatedAt int64              `bson:"createdAt" json:"createdAt"`
}
