package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	userService *data.UserService
	jwtSecret   string
}

func NewAuthController(userService *data.UserService, jwtSecret string) *AuthController {
	return &AuthController{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := ac.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID.Hex(),
		"username": createdUser.Username,
		"role":     createdUser.Role,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userService.VerifyUser(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
	})

	tokenString, err := token.SignedString([]byte(ac.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return token in both body and header for convenience
	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func (ac *AuthController) Logout(c *gin.Context) {
	// For simplicity, logout can be handled by client-side token deletion.
	// In a real application, you might want to implement token invalidation.
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}