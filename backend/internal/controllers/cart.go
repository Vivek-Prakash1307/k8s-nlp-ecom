package controllers

import (
	"context"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddToCart(c *gin.Context) {
	userIDraw, _ := c.Get("user_id")
	userIDHex := userIDraw.(string)
	uid, _ := primitive.ObjectIDFromHex(userIDHex)

	var body struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	pid, _ := primitive.ObjectIDFromHex(body.ProductID)

	client := db.Connect()
	coll := client.Database(os.Getenv("DB_NAME")).Collection("carts")
	ctx := context.Background()

	// upsert: push or increase quantity
	var cart models.Cart
	err := coll.FindOne(ctx, bson.M{"userId": uid}).Decode(&cart)
	if err != nil {
		// create
		newCart := models.Cart{
			ID:        primitive.NewObjectID(),
			UserID:    uid,
			Items:     []models.CartItem{{ProductID: pid, Quantity: body.Quantity}},
			UpdatedAt: time.Now().Unix(),
		}
		coll.InsertOne(ctx, newCart)
		c.JSON(http.StatusOK, gin.H{"cart": newCart})
		return
	}
	// update existing
	updated := false
	for i := range cart.Items {
		if cart.Items[i].ProductID == pid {
			cart.Items[i].Quantity += body.Quantity
			updated = true
			break
		}
	}
	if !updated {
		cart.Items = append(cart.Items, models.CartItem{ProductID: pid, Quantity: body.Quantity})
	}
	cart.UpdatedAt = time.Now().Unix()
	coll.ReplaceOne(ctx, bson.M{"_id": cart.ID}, cart)
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func GetCart(c *gin.Context) {
	userIDraw, _ := c.Get("user_id")
	userIDHex := userIDraw.(string)
	uid, _ := primitive.ObjectIDFromHex(userIDHex)
	client := db.Connect()
	coll := client.Database(os.Getenv("DB_NAME")).Collection("carts")
	ctx := context.Background()
	var cart models.Cart
	if err := coll.FindOne(ctx, bson.M{"userId": uid}).Decode(&cart); err != nil {
		c.JSON(http.StatusOK, gin.H{"cart": models.Cart{Items: []models.CartItem{}}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}
