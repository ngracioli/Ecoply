package middlewares

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"ecoply/internal/domain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupplierMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User = utils.GetUserFromContext(c)
		var userTypeService services.UserTypeService = services.UserType

		if !userTypeService.UserIsSupplier(user) {
			err := merr.NewResponseError(http.StatusForbidden, ErrUserIsNotSupplier)
			c.AbortWithStatusJSON(err.StatusCode, err)
			return
		}

		c.Next()
	}
}
