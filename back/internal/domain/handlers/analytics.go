package handlers

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandlers interface {
	Platform(c *gin.Context)
	User(c *gin.Context)
}

type analyticsHandlers struct {
	analyticsServices services.AnalyticsService
}

func NewAnalyticsHandler(analyticsServices services.AnalyticsService) AnalyticsHandlers {
	return &analyticsHandlers{
		analyticsServices: analyticsServices,
	}
}

func (h *analyticsHandlers) Platform(c *gin.Context) {
	response, err := h.analyticsServices.Platform()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *analyticsHandlers) User(c *gin.Context) {
	var user *models.User = GetUserFromContext(c)
	response, err := h.analyticsServices.User(user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
