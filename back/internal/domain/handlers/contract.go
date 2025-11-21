package handlers

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractHandlers interface {
	Get(c *gin.Context)
}

type contractHandlers struct {
	contractService services.ContractService
}

func NewContractHandlers(service services.ContractService) ContractHandlers {
	return &contractHandlers{
		contractService: service,
	}
}

func (h *contractHandlers) Get(c *gin.Context) {
	var purchaseUuid string = c.Param("uuid")
	var user *models.User = GetUserFromContext(c)

	response, err := h.contractService.Get(user, purchaseUuid)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
