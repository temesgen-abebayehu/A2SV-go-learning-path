package router

import (
	"task_manager/controllers"
	"task_manager/data"
	// "task_manager/middleware"
	// "task_manager/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(secretKey string) *gin.Engine {
	r := gin.Default()

	// Initialize services
	userService := data.NewUserService()

	// Initialize controllers
	authController := controllers.NewAuthController(userService, secretKey)

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
		api.POST("/logout", authController.Logout)
		
	}

	return r
}