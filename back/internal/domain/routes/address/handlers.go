package address

import (
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddressByCepHandler(c *gin.Context) {
	cep := c.Param("cep")

	addressData, err := services.LoadAddressByCep(cep)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	response := addressDataToResource(addressData)

	c.JSON(http.StatusOK, gin.H{"data": response})
}
