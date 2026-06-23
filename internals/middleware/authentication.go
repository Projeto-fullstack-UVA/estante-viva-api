package middleware

import (
	"os"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		c.AbortWithStatus(500)
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(401)
		return
	}

	tokenString = auth.RemoveBearerPrefix(tokenString)
	if tokenString == "" {
		c.AbortWithStatus(401)
		return
	}

	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	c.Set("user_id", claims["user_id"])
	c.Set("role", claims["role"])
	c.Next()
}
