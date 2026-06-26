package middleware

import (
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(c *gin.Context) {
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

func GetCurrentUserID(c *gin.Context) (int64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	switch id := userID.(type) {
	case float64:
		return int64(id), true
	case int64:
		return id, true
	}
	return 0, false
}
