package main

import (
	"os"
	"time"

	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	r := gin.Default()
	// connect db
	db.Connect()
	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_ORIGIN")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Setup(r)
	r.Run(":8080")
}
