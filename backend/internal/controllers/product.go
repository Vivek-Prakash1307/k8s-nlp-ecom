package controllers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListProducts(c *gin.Context) {
	q := c.Query("q")
	sortBy := c.DefaultQuery("sort", "relevance") // relevance, price_asc, price_desc, rating, popularity
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	limit := 20
	skip := (page - 1) * limit

	client := db.Connect()
	coll := client.Database(os.Getenv("DB_NAME")).Collection("products")

	filter := bson.M{}
	if q != "" {
		filter = bson.M{"$text": bson.M{"$search": q}}
	}
	opts := options.Find()
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))
	switch sortBy {
	case "price_asc":
		opts.SetSort(bson.M{"price": 1})
	case "price_desc":
		opts.SetSort(bson.M{"price": -1})
	case "rating":
		opts.SetSort(bson.M{"rating": -1})
	case "popularity":
		opts.SetSort(bson.M{"popularity": -1})
	default:
		opts.SetSort(bson.M{"createdAt": -1})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := coll.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db"})
		return
	}
	var res []models.Product
	if err := cur.All(ctx, &res); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db decode"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": res})
}
