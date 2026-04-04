package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "No role information found",
			})
			c.Abort()
			return
		}

		userRole := role.(string)
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "Insufficient permissions",
		})
		c.Abort()
	}
}

// Helper to get current user from context
func GetCurrentUser(c *gin.Context) (uint, string, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, "", errors.New("user not authenticated")
	}

	email, exists := c.Get("email")
	if !exists {
		return 0, "", errors.New("email not found in context")
	}

	return userID.(uint), email.(string), nil
}
