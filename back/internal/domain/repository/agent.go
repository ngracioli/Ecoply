package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type AgentRepository interface {
	WithTransaction(tx *gorm.DB) AgentRepository

	Create(agent *models.Agent) (*models.Agent, error)
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

func (a *agentRepository) WithTransaction(tx *gorm.DB) AgentRepository {
	return NewAgentRepository(tx)
}

func (a *agentRepository) Create(agent *models.Agent) (*models.Agent, error) {
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
