package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type AgentRepository interface {
	Create(params AgentCreateParams) (*models.Agent, error)
	FindById(id uint) (*models.Agent, error)
	FindByCnpj(cnpj string) (*models.Agent, error)
	FindByCceeCode(cceeCode string) (*models.Agent, error)
}

type agentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &agentRepository{db: db}
}

type AgentCreateParams struct {
	Cnpj        string
	CompanyName string
	CceeCode    string
	Submarket   *models.Submarket
	Address     *models.Address
}

func (a *agentRepository) Create(params AgentCreateParams) (*models.Agent, error) {
	agent := &models.Agent{
		Cnpj:        params.Cnpj,
		CompanyName: params.CompanyName,
		CceeCode:    params.CceeCode,
		SubmarketId: params.Submarket.ID,
		AddressId:   params.Address.ID,
	}

	if err := a.db.Create(agent).Error; err != nil {
		mlog.Log("Failed to create agent: " + err.Error())
		return nil, err
	}

	return agent, nil
}

func (a *agentRepository) FindById(id uint) (*models.Agent, error) {
	var agent models.Agent
	err := a.db.First(&agent, id).Error

	if err != nil {
		mlog.Log("Failed to find agent by ID: " + err.Error())
		return nil, err
	}

	return &agent, nil
}

func (a *agentRepository) FindByCnpj(cnpj string) (*models.Agent, error) {
	var agent models.Agent
	err := a.db.Where("cnpj = ?", cnpj).First(&agent).Error

	if err != nil {
		mlog.Log("Failed to find agent by CNPJ: " + err.Error())
		return nil, err
	}

	return &agent, nil
}

func (a *agentRepository) FindByCceeCode(cceeCode string) (*models.Agent, error) {
	var agent models.Agent
	err := a.db.Where("ccee_code = ?", cceeCode).First(&agent).Error

	if err != nil {
		mlog.Log("Failed to find agent by CCEE Code: " + err.Error())
		return nil, err
	}

	return &agent, nil
}
