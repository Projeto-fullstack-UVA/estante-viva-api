package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/controllers"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/middleware"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var allowedOrigins []string

func getAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		log.Fatalln("Environment variable ALLOWED_ORIGINS is not set")
	}

	parts := strings.Split(origins, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	log.Println("Loaded origins variables with success")
	return parts
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allowOrigin := allowedOrigins[0]
		if origin != "" && slices.Contains(allowedOrigins, origin) {
			allowOrigin = origin
		}

		c.Header("Access-Control-Allow-Origin", allowOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
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
		log.Fatalln("No .env file found")
	}

	jwtKey := os.Getenv("JWT_SECRET_KEY")
	if jwtKey == "" {
		log.Fatalln("Environment variable JWT_SECRET_KEY is not set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalln("Environment variable DATABASE_URL is not set")
	}

	allowedOrigins = getAllowedOrigins()

	log.Println("Success loading the environment variables")

	if err := repositories.Init(databaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected")

	router := gin.Default()
	router.Use(cors())

	router.POST("/login", controllers.Login)
	router.GET("/users", middleware.Authentication, controllers.ListUsers)
	router.POST("/users", controllers.Register)
	router.GET("/users/:id", middleware.Authentication, controllers.FindUser)
	router.PATCH("/users/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.UpdateUser)
	router.DELETE("/users/:id", middleware.Authentication, controllers.DeleteUser)
	router.GET("/books", middleware.Authentication, controllers.ListBooks)
	router.POST("/books", middleware.Authentication, controllers.CreateBook)
	router.GET("/books/:id", middleware.Authentication, controllers.FindBook)
	router.PATCH("/books/:id", middleware.Authentication, controllers.UpdateBook)
	router.DELETE("/books/:id", middleware.Authentication, controllers.DeleteBook)
	router.GET("/loans", middleware.Authentication, controllers.ListLoans)
	router.POST("/loans", middleware.Authentication, controllers.BorrowBook)
	router.GET("/loans/:id", middleware.Authentication, controllers.FindLoan)
	router.PATCH("/loans/:id", middleware.Authentication, controllers.ReturnBook)
	router.DELETE("/loans/:id", middleware.Authentication, controllers.DeleteLoan)

	log.Println("Server running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
