package address

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CompanyByCnpj(c *gin.Context) {
	cnpj := c.Param("cnpj")

	companyData, err := services.LoadCnpjData(cnpj)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	response := companyDataToResource(companyData)

	c.JSON(http.StatusOK, gin.H{"data": response})
}
