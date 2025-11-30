package handlers

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CceeHandlers interface {
	GetAgentsByCnpj(c *gin.Context)
}

type cceeHandler struct {
	cceeService services.CceeService
}

func NewCceeHandler(cceeService services.CceeService) CceeHandlers {
	return &cceeHandler{
		cceeService: cceeService,
	}
}

func (h *cceeHandler) GetAgentsByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	agents, err := h.cceeService.GetAgentsByCnpj(cnpj)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": agents})
}
