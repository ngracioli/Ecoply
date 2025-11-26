package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/resources"
	"ecoply/internal/domain/services"
	"ecoply/internal/domain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseHandlers interface {
	Create(c *gin.Context)
	ListPurchases(c *gin.Context)
	ListSales(c *gin.Context)
	FindByUuid(c *gin.Context)
	Cancel(c *gin.Context)
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

	response, err := h.purchaseService.Create(&payload, offerUuid, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (h *purchaseHandlers) ListPurchases(c *gin.Context) {
	var response *utils.PaginationWrapper[*resources.Purchase]
	var payload requests.ListPurchase
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.purchaseService.ListPurchases(&payload, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *purchaseHandlers) ListSales(c *gin.Context) {
	var response *utils.PaginationWrapper[*resources.Purchase]
	var payload requests.ListSold
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.purchaseService.ListSold(&payload, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *purchaseHandlers) Cancel(c *gin.Context) {
	var purchaseUuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	err := h.purchaseService.Cancel(purchaseUuid, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func (h *purchaseHandlers) FindByUuid(c *gin.Context) {
	var purchaseUuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	response, err := h.purchaseService.FindByUuid(user, purchaseUuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
