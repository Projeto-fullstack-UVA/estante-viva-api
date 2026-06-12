package main

import (
	"log"
	"net/http"
	"os"
	"slices"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/controllers"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var allowedOrigins = []string{"http://localhost:5173", "http://localhost:4173"}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allowOrigin := allowedOrigins[0]
		if origin != "" && slices.Contains(allowedOrigins, origin) {
			allowOrigin = origin
		}

		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Vary", "Origin")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	if err := repositories.Init(databaseURL); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := gin.Default()
	r.Use(cors())

	r.POST("/login", controllers.Login)
	r.GET("/users", controllers.ListUsers)
	r.POST("/users", controllers.Register)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/books", controllers.ListBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.GET("/loans", controllers.ListLoans)
	r.POST("/loans", controllers.BorrowBook)
	r.GET("/loans/:id", controllers.FindLoan)
	r.PATCH("/loans/:id", controllers.ReturnBook)

	log.Println("Server running on http://localhost:3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
