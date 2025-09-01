package middlewares

import (
	"ecoply/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO revisaar isso aqui, refatorar os maps puros de erro para a struct de erro e trocar as mensagens por constantes
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := services.ExtractTokenFromBearer(authHeader)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		jwtService := services.NewJWTService()
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_type", claims.UserType)
		c.Set("claims", claims)

		c.Next()
	}
}
