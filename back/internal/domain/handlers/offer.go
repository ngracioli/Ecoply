package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OfferByUuidHanlder(c *gin.Context) {
	var uuid string = c.Param("uuid")

	response, err := services.Offer.GetByUuid(uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func OfferListHandler(c *gin.Context) {
	var params requests.ListOffers
	var user *models.User

	if err := c.ShouldBindQuery(&params); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	user = GetUserFromContext(c)
	response, err := services.Offer.List(&params, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UserOffersHandler(c *gin.Context) {
	var user *models.User = GetUserFromContext(c)

	response, err := services.Offer.BelongingToUser(user.ID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func CreateOfferHandler(c *gin.Context) {
	var payload requests.CreateOffer
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	var user *models.User = GetUserFromContext(c)
	response, err := services.Offer.Create(user, &payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func UpdateOfferHandler(c *gin.Context) {
	var payload requests.UpdateOffer
	var uuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	err := services.Offer.Update(user, uuid, &payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func DeleteOfferHandler(c *gin.Context) {
	var uuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	err := services.Offer.Delete(user, uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
