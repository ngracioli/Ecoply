package middlewares

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var responseError *merr.ResponseError
		defer func() {
			if responseError != nil {
				c.JSON(responseError.StatusCode, responseError)
				c.Abort()
			}
		}()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responseError = merr.NewResponseError(http.StatusUnauthorized, ErrJwtAuthorizationHeaderRequired)
			return
		}

		tokenString := extractBearerToken(authHeader)
		if tokenString == "" {
			responseError = merr.NewResponseError(http.StatusUnauthorized, ErrJwtBearerTokenRequired)
			return
		}

		jwtService := services.NewJwtService()
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			responseError = merr.NewResponseError(http.StatusUnauthorized, ErrJwtInvalidToken)
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_type", claims.UserType)
		c.Set("claims", claims)

		c.Next()
	}
}

func extractBearerToken(header string) string {
	var slice []string = strings.Split(header, "Bearer ")
	if len(slice) == 2 {
		return slice[1]
	}
	return ""
}
