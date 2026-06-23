package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"
)

// Authorization restricts access to handlers to the given roles.
// It must run after Authentication, which sets the "role" context value
// from the verified token claims.
func Authorization(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("role")
		if !exists {
			c.AbortWithStatus(401)
			return
		}

		role, ok := value.(string)
		if !ok || role == "" {
			c.AbortWithStatus(401)
			return
		}

		if !slices.Contains(roles, role) {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
