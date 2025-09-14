package handlers

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CompanyByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	response, err := services.LoadCnpjData(cnpj)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
