package middleware

import (
	"net/http"
	"strings"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(requiredRole models.UserRole, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
				"details": "Format should be: 'Bearer <token>'",
			})
			return
		}

		// 2. Check Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
				"details": "Format should be: 'Bearer <token>'",
			})
			return
		}

		tokenString := parts[1]

		// 3. Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
				"details": err.Error(),
			})
			return
		}

		// 4. Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		// 5. Role check (if required)
		if requiredRole != "" {
			role, ok := claims["role"].(string)
			if !ok || models.UserRole(role) != requiredRole {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": "Insufficient permissions",
					"details": "Required role: " + string(requiredRole),
				})
				return
			}
		}

		// 6. Set user info in context
		c.Set("userID", claims["sub"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()
	}
}