package environment

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var AllowedOrigins []string
var JwtSecretKey string
var DatabaseURL string
var Port string
var GinMode string

func LoadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	AllowedOrigins = getAllowedOrigins()

	JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	if JwtSecretKey == "" {
		log.Fatalln("Environment variable JWT_SECRET_KEY is not set")
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		log.Fatalln("Environment variable DATABASE_URL is not set")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		log.Fatalln("Environment variable PORT is not set")
	}

	GinMode = os.Getenv("GIN_MODE")
	if GinMode == "" {
		log.Fatalln("Environment variable GIN_MODE is not set")
	}

	log.Println("Success loading the environment variables")
}

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
