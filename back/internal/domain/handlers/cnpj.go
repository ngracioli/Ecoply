package handlers

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CnpjHandlers interface {
	CompanyByCnpj(c *gin.Context)
}

type cnpjHandler struct{}

func NewCnpjHandler() CnpjHandlers {
	return &cnpjHandler{}
}

func (h *cnpjHandler) CompanyByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	response, err := services.LoadCnpjData(cnpj)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
