package main

import (
	"os"
	"strings"
	"time"

	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/db"
	"github.com/Vivek-Prakash1307/k8s-nlp-ecom/backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	// Connect DB
	db.Connect()

	// CORS configuration
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

	// Safety fallback if env missing or wrong formatted
	if !strings.HasPrefix(frontendOrigin, "http://") &&
		!strings.HasPrefix(frontendOrigin, "https://") {
		frontendOrigin = "http://localhost:3000"
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.Setup(r)
	r.Run(":8080")
}
