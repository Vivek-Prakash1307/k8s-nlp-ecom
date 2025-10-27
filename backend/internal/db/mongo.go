package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	once   sync.Once
)

func Connect() *mongo.Client {
	once.Do(func() {
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			log.Fatal("MONGO_URI not set")
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
		Client = client
	})
	return Client
}
