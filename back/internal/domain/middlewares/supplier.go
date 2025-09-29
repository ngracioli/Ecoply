package middlewares

import (
	"ecoply/internal/domain/merr"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupplierMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := c.GetString("user_type")

		if userType != "supplier" {
			err := merr.NewResponseError(http.StatusForbidden, ErrUserIsNotSupplier)
			c.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		c.Next()
	}
}
