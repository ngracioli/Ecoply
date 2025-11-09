package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePurchaseHandler(c *gin.Context) {
	var payload requests.CreatePurchase
	var offerUuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	err := services.Purchase.Create(&payload, offerUuid, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
