package controllers

import (
    "context"
    "net/http"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
    "github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/models"
)

func Checkout(c *gin.Context) {
    userIDraw, _ := c.Get("user_id")
    userIDHex := userIDraw.(string)
    uid, _ := primitive.ObjectIDFromHex(userIDHex)

    var body struct{ Address string `json:"address"` }
    if err := c.BindJSON(&body); err!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid"}); return }

    client := db.Connect()
    cartCol := client.Database(os.Getenv("DB_NAME")).Collection("carts")
    prodCol := client.Database(os.Getenv("DB_NAME")).Collection("products")
    orderCol := client.Database(os.Getenv("DB_NAME")).Collection("orders")

    ctx := context.Background()
    var cart models.Cart
    if err := cartCol.FindOne(ctx, bson.M{"userId": uid}).Decode(&cart); err!=nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"empty cart"})
        return
    }

    // compute total
    total := 0.0
    for _, it := range cart.Items {
        var prod models.Product
        if err := prodCol.FindOne(ctx, bson.M{"_id": it.ProductID}).Decode(&prod); err==nil {
            total += prod.Price * float64(it.Quantity)
        }
    }

    order := models.Order{
        ID: primitive.NewObjectID(),
        UserID: uid,
        Items: cart.Items,
        Total: total,
        Status: "confirmed", // just for the temporary , when we install particualarly for the actaul payment gateway we will integrate that part 
        Address: body.Address,
        CreatedAt: time.Now().Unix(),
    }

    if _, err := orderCol.InsertOne(ctx, order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error":"order failed"})
        return
    }
    // clear cart
    cartCol.DeleteOne(ctx, bson.M{"_id": cart.ID})
    c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
    userIDraw, _ := c.Get("user_id")
    userIDHex := userIDraw.(string)
    uid, _ := primitive.ObjectIDFromHex(userIDHex)
    client := db.Connect()
    coll := client.Database(os.Getenv("DB_NAME")).Collection("orders")
    ctx := context.Background()
    cur, _ := coll.Find(ctx, bson.M{"userId": uid})
    var orders []models.Order
    cur.All(ctx, &orders)
    c.JSON(http.StatusOK, gin.H{"orders": orders})
}
