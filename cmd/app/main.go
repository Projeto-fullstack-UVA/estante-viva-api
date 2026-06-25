package main

import (
	"log"
	"net/http"
	"slices"

	environment "github.com/Projeto-fullstack-UVA/estante-viva-api/internals"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/controllers"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/middleware"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allowOrigin := environment.AllowedOrigins[0]
		if origin != "" && slices.Contains(environment.AllowedOrigins, origin) {
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
	environment.LoadEnvironmentVariables()
	if err := repositories.Init(environment.DatabaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established with success")

	router := gin.Default()
	router.Use(cors())

	router.GET("/me", middleware.Authentication, controllers.GetMe)
	router.GET("/users", middleware.Authentication, middleware.Authorization("admin"), controllers.ListUsers)
	router.GET("/users/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.FindUser)
	router.GET("/books", middleware.Authentication, controllers.ListBooks)
	router.GET("/books/:id", middleware.Authentication, controllers.FindBook)
	router.GET("/loans", middleware.Authentication, controllers.ListLoans)
	router.GET("/loans/:id", middleware.Authentication, controllers.FindLoan)
	router.GET("/institutions", controllers.ListInstitutions)
	router.GET("/institutions/:id", middleware.Authentication, controllers.FindInstitution)
	router.GET("/events", middleware.Authentication, controllers.ListEvents)
	router.GET("/events/:id", middleware.Authentication, controllers.FindEvent)
	router.POST("/login", controllers.Login)
	router.POST("/users", controllers.Register)
	router.POST("/books", middleware.Authentication, middleware.Authorization("admin", "teacher"), controllers.CreateBook)
	router.POST("/loans", middleware.Authentication, controllers.BorrowBook)
	router.POST("/institutions", middleware.Authentication, middleware.Authorization("admin"), controllers.CreateInstitution)
	router.POST("/events", middleware.Authentication, middleware.Authorization("admin", "teacher"), controllers.CreateEvent)
	router.PATCH("/users/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.UpdateUser)
	router.PATCH("/books/:id", middleware.Authentication, middleware.Authorization("admin", "teacher"), controllers.UpdateBook)
	router.PATCH("/loans/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.ReturnBook)
	router.DELETE("/users/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.DeleteUser)
	router.DELETE("/books/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.DeleteBook)
	router.DELETE("/loans/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.DeleteLoan)
	router.DELETE("/institutions/:id", middleware.Authentication, middleware.Authorization("admin"), controllers.DeleteInstitution)
	router.DELETE("/events/:id", middleware.Authentication, middleware.Authorization("admin", "teacher"), controllers.DeleteEvent)

	if err := router.Run(environment.Port); err != nil {
		log.Fatalln(err)
	}
}
