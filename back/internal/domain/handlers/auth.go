package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var payload requests.Login

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := services.Auth.Login(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func SignUpHandler(c *gin.Context) {
	var payload requests.SignUp

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := services.Auth.SignUp(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func MeHandler(c *gin.Context) {
	var user *models.User = GetUserFromContext(c)
	response, err := services.Auth.Me(user.Uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func AvailabilityHandler(c *gin.Context) {
	var payload requests.Availability

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := services.Auth.Availability(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	if !response {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func RefreshTokenHandler(c *gin.Context) {
	token := c.GetString("token")
	response, err := services.Auth.RefreshToken(token)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": response}})
}
