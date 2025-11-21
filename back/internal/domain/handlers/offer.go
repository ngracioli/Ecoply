package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OfferHandlers interface {
	FindByUuid(c *gin.Context)
	List(c *gin.Context)
	FromUser(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Purchases(c *gin.Context)
}

type offerHandler struct {
	offerService services.OfferService
}

func NewOfferHandler(offerService services.OfferService) OfferHandlers {
	return &offerHandler{
		offerService: offerService,
	}
}

func (h *offerHandler) FindByUuid(c *gin.Context) {
	var uuid string = c.Param("uuid")

	response, err := h.offerService.GetByUuid(uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *offerHandler) List(c *gin.Context) {
	var params requests.ListOffers
	var user *models.User

	if err := c.ShouldBindQuery(&params); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	user = GetUserFromContext(c)
	response, err := h.offerService.List(&params, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *offerHandler) FromUser(c *gin.Context) {
	var user *models.User = GetUserFromContext(c)

	response, err := h.offerService.BelongingToUser(user.ID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *offerHandler) Create(c *gin.Context) {
	var payload requests.CreateOffer
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	var user *models.User = GetUserFromContext(c)
	response, err := h.offerService.Create(user, &payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (h *offerHandler) Update(c *gin.Context) {
	var payload requests.UpdateOffer
	var uuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	err := h.offerService.Update(user, uuid, &payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func (h *offerHandler) Delete(c *gin.Context) {
	var uuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	err := h.offerService.Delete(user, uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func (h *offerHandler) Purchases(c *gin.Context) {
	var params requests.ListPurchasesFromOffer
	var uuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindQuery(&params); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.offerService.Purchases(uuid, &params, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
