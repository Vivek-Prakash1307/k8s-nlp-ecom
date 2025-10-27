package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/controllers"
    "github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/middleware"
)

func Setup(r *gin.Engine) {
    r.POST("/api/signup", controllers.Signup)
    r.POST("/api/login", controllers.Login)

    r.GET("/api/products", controllers.ListProducts)

    auth := r.Group("/api")
    auth.Use(middleware.AuthRequired())
    {
        auth.POST("/cart/add", controllers.AddToCart)
        auth.GET("/cart", controllers.GetCart)
        auth.POST("/checkout", controllers.Checkout)
        auth.GET("/orders", controllers.GetOrders)
    }
}
