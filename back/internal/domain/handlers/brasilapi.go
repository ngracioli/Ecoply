package handlers

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrasilApiHandlers interface {
	GetStates(c *gin.Context)
	GetCnpjData(c *gin.Context)
	GetCepData(c *gin.Context)
}

type brasilApiHandler struct {
	brasilApiService services.BrasilApiService
}

func NewBrasilApiHandler(brasilApiService services.BrasilApiService) BrasilApiHandlers {
	return &brasilApiHandler{
		brasilApiService: brasilApiService,
	}
}

func (h *brasilApiHandler) GetStates(c *gin.Context) {
	states, err := h.brasilApiService.GetStates()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": states})
}

func (h *brasilApiHandler) GetCnpjData(c *gin.Context) {
	cnpj := c.Param("cnpj")

	cnpjData, err := h.brasilApiService.GetCnpjData(cnpj)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cnpjData})
}

func (h *brasilApiHandler) GetCepData(c *gin.Context) {
	cep := c.Param("cep")

	cepData, err := h.brasilApiService.GetCepData(cep)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cepData})
}
