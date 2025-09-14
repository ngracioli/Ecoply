package handlers

import (
	"ecoply/internal/database"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService services.AuthService

func GetAuthService() services.AuthService {
	if authService == nil {
		authService = services.NewAuthService(database.Con)
	}
	return authService
}

func LoginHandler(c *gin.Context) {
	var payload requests.Login

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := GetAuthService().Login(&payload)
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

	response, err := GetAuthService().SignUp(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func MeHandler(c *gin.Context) {
	userUuid, _ := c.Get("user_uuid")
	response, err := GetAuthService().Me(userUuid.(string))
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

	response, err := GetAuthService().Availability(&payload)
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
	response, err := GetAuthService().RefreshToken(token)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": response}})
}
