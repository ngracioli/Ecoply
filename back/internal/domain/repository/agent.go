package repository

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type AgentRepository interface {
	Create(params AgentCreateParams) (*models.Agent, *merr.ResponseError)
	FindById(id uint) (*models.Agent, *merr.ResponseError)
	FindByCnpj(cnpj string) (*models.Agent, *merr.ResponseError)
	FindByCceeCode(cceeCode string) (*models.Agent, *merr.ResponseError)
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

func (a *agentRepository) Create(params AgentCreateParams) (*models.Agent, *merr.ResponseError) {
	existing, err := a.FindByCnpj(params.Cnpj)
	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	if existing != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrAgentCnpjAlreadyExists)
	}

	existing, err = a.FindByCceeCode(params.CceeCode)
	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	if existing != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrAgentCceeCodeAlreadyExists)
	}

	agent := &models.Agent{
		Cnpj:        params.Cnpj,
		CompanyName: params.CompanyName,
		CceeCode:    params.CceeCode,
		SubmarketId: params.Submarket.ID,
		AddressId:   params.Address.ID,
	}

	if err := a.db.Create(agent).Error; err != nil {
		mlog.Log("Failed to create agent: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return agent, nil
}

func (a *agentRepository) FindById(id uint) (*models.Agent, *merr.ResponseError) {
	var agent models.Agent
	err := a.db.First(&agent, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find agent by ID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &agent, nil
}

func (a *agentRepository) FindByCnpj(cnpj string) (*models.Agent, *merr.ResponseError) {
	var agent models.Agent
	err := a.db.Where("cnpj = ?", cnpj).First(&agent).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find agent by CNPJ: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &agent, nil
}

func (a *agentRepository) FindByCceeCode(cceeCode string) (*models.Agent, *merr.ResponseError) {
	var agent models.Agent
	err := a.db.Where("ccee_code = ?", cceeCode).First(&agent).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find agent by CCEE Code: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &agent, nil
}
