package handlers

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlers interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
	Me(c *gin.Context)
	Availability(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authHandlers struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandlers {
	return &authHandlers{
		authService: authService,
	}
}

func (h *authHandlers) Login(c *gin.Context) {
	var payload requests.Login

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.authService.Login(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *authHandlers) SignUp(c *gin.Context) {
	var payload requests.SignUp

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.authService.SignUp(&payload)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": response})
}

func (h *authHandlers) Me(c *gin.Context) {
	var user *models.User = GetUserFromContext(c)
	response, err := h.authService.Me(user.Uuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *authHandlers) Availability(c *gin.Context) {
	var payload requests.Availability

	if err := c.ShouldBindQuery(&payload); err != nil {
		var response *merr.ResponseError = merr.BindValidationErrorsToResponse(err)
		c.JSON(response.StatusCode, response)
		return
	}

	response, err := h.authService.Availability(&payload)
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

func (h *authHandlers) RefreshToken(c *gin.Context) {
	token := c.GetString("token")
	response, err := h.authService.RefreshToken(token)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": response}})
}
