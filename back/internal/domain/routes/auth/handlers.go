package auth

import (
	"ecoply/internal/domain/merr"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var payload LoginRequest

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := Login(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func SignUpHandler(c *gin.Context) {
	var payload SignUpRequest

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := SignUp(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func MeHandler(c *gin.Context) {
	userUuid, _ := c.Get("user_uuid")
	response, err := Me(userUuid.(string))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func IsEmailAvailableHandler(c *gin.Context) {
	var payload IsEmailAvailableRequest

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := IsEmailAvailable(&payload)
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

func IsCpfCnpjAvailableHandler(c *gin.Context) {
	var payload IsCpfCnpjAvailableRequest

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := IsCpfCnpjAvailable(&payload)
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
