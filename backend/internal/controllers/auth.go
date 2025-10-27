package controllers

import (
	"context"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/models"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct{ Name, Email, Password string }
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	client := db.Connect()
	coll := client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx := context.Background()
	// check existing
	var existing models.User
	if err := coll.FindOne(ctx, bson.M{"email": body.Email}).Decode(&existing); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email exists"})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	u := models.User{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Email:     body.Email,
		Password:  string(hash),
		CreatedAt: time.Now().Unix(),
	}
	if _, err := coll.InsertOne(ctx, u); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	token, _ := utils.GenerateToken(u.ID.Hex())
	c.JSON(http.StatusOK, gin.H{"token": token, "user": gin.H{"id": u.ID.Hex(), "name": u.Name, "email": u.Email}})
}

func Login(c *gin.Context) {
	var body struct{ Email, Password string }
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	client := db.Connect()
	coll := client.Database(os.Getenv("DB_NAME")).Collection("users")
	ctx := context.Background()
	var u models.User
	if err := coll.FindOne(ctx, bson.M{"email": body.Email}).Decode(&u); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, _ := utils.GenerateToken(u.ID.Hex())
	c.JSON(http.StatusOK, gin.H{"token": token, "user": gin.H{"id": u.ID.Hex(), "name": u.Name, "email": u.Email}})
}
