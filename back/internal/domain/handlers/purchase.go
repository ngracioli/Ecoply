package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseHandlers interface {
	Create(c *gin.Context)
}

type purchaseHandlers struct {
	purchaseService services.PurchaseService
}

func NewPurchaseHandlers(purchaseService services.PurchaseService) PurchaseHandlers {
	return &purchaseHandlers{
		purchaseService: purchaseService,
	}
}

func (h *purchaseHandlers) Create(c *gin.Context) {
	var payload requests.CreatePurchase
	var offerUuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	err := h.purchaseService.Create(&payload, offerUuid, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
