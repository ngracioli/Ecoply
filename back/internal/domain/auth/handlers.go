package auth

import (
	"ecoply/internal/domain/merr"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var payload LoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := login(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
	}

	c.JSON(200, response)
}

func SignUpHandler(c *gin.Context) {
	var payload SignUpRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	// response, err := signUp(&payload)

	c.JSON(201, gin.H{})
}

func MeHandler(c *gin.Context) {

}
