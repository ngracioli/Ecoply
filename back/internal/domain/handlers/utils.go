package handlers

import (
	"ecoply/internal/domain/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) *models.User {
	user, _ := c.Get("user")
	return user.(*models.User)
}
