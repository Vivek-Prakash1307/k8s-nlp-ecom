package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	godotenv.Load()
	client := db.Connect()
	coll := client.Database("ecomdb").Collection("products")
	ctx := context.Background()
	coll.Drop(ctx)
	rand.Seed(time.Now().UnixNano())
	cats := []string{"electronics", "books", "clothing", "home", "sports", "beauty"}
	for i := 1; i <= 150; i++ {
		p := models.Product{
			ID:          primitive.NewObjectID(),
			Title:       fmt.Sprintf("Product %03d", i),
			Description: "Dummy product for testing",
			Price:       float64(rand.Intn(5000))/10.0 + 5.0,
			Rating:      float64((rand.Intn(40) + 10)) / 10.0,
			Popularity:  rand.Intn(1000),
			Category:    cats[rand.Intn(len(cats))],
			CreatedAt:   time.Now().Unix(),
		}
		coll.InsertOne(ctx, p)
	}
	fmt.Println("Seeded 150 products")
}
