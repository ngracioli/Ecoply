package middlewares

import (
	"ecoply/internal/domain/handlers"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupplierMiddleware(userTypeService services.UserTypeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User = handlers.GetUserFromContext(c)

		if !userTypeService.UserIsSupplier(user) {
			err := merr.NewResponseError(http.StatusForbidden, ErrUserIsNotSupplier)
			c.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		c.Next()
	}
}
