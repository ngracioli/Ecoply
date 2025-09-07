package middlewares

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/services"
	"errors"
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

		c.Set("user_uuid", claims.UserUuid)
		c.Set("user_email", claims.UserEmail)
		c.Set("user_type", claims.UserType)
		c.Set("claims", claims)

		if err := ensureClaimsExists(c); err != nil {
			responseError = err
			return
		}

		c.Next()
	}
}

func ensureClaimsExists(c *gin.Context) *merr.ResponseError {
	if err := ensureClaim[string](c, "user_uuid"); err != nil {
		return err
	}
	if err := ensureClaim[string](c, "user_email"); err != nil {
		return err
	}
	if err := ensureClaim[uint](c, "user_type"); err != nil {
		return err
	}
	return nil
}

func ensureClaim[T uint | string](c *gin.Context, key string) *merr.ResponseError {
	claim, exists := c.Get(key)

	if !exists {
		return merr.NewResponseError(http.StatusForbidden, errors.Join(ErrMissingClaim, errors.New(" "+key)))
	}

	var valid bool = false

	if v, ok := claim.(string); ok {
		if v != "" {
			valid = true
		}
	}

	if v, ok := claim.(uint); ok {
		if v != 0 {
			valid = true
		}
	}

	if !valid {
		return merr.NewResponseError(http.StatusForbidden, errors.Join(ErrMissingClaim, errors.New(" "+key)))
	}

	return nil
}

func extractBearerToken(header string) string {
	var slice []string = strings.Split(header, "Bearer ")
	if len(slice) == 2 {
		return slice[1]
	}
	return ""
}
