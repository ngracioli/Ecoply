package middlewares

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
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

		var user *models.User
		user, responseError = services.User.FindByUuid(claims.UserUuid)
		if responseError != nil {
			return
		}

		c.Set("user", user)
		c.Set("token", tokenString)
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
