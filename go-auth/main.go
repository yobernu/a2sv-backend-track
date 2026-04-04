package main

import (
	"go-auth/controllers"
	"go-auth/models"
	"go-auth/services"
	"go-auth/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize dependencies
	userStore := models.NewUserStore()
	userService := services.NewUserService(userStore)
	userController := controllers.NewUserController(userService)

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Authentication API",
			"version": "1.0.0",
		})
	})

	router.POST("/register", userController.Register)
	router.POST("/login", userController.LoginHandler)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Protected routes requiring authentication
	authenticated := router.Group("/")
	authenticated.Use(utils.AuthMiddleware())
	{
		authenticated.GET("/profile", profileHandler)
		// authenticated.PUT("/profile", updateProfileHandler)
	}

	// Admin-only routes
	admin := router.Group("/admin")
	admin.Use(utils.AuthMiddleware(), utils.RequireRole("admin"))
	{
		admin.GET("/users", listUsersHandler)
		// admin.DELETE("/users/:id", deleteUserHandler)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func profileHandler(c *gin.Context) {
	userID, email, err := utils.GetCurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"email":   email,
	})
}

func listUsersHandler(c *gin.Context) {
	// Admin-only functionality
	c.JSON(http.StatusOK, gin.H{
		"message": "List of all users (admin only)",
	})
}
