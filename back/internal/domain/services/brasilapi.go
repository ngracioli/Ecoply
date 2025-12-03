package services

import (
	"context"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/resources"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrFailedToFetchBrasilApi = errors.New("Erro ao buscar dados. Tente novamente.")
	ErrBrasilApiNotFound      = errors.New("Dados não encontrados")
	ErrFailedToParseBrasilApi = errors.New("Erro ao processar dados. Tente novamente.")
	ErrInvalidCepFormat       = errors.New("CEP inválido")
	ErrCepNotFound            = errors.New("CEP não encontrado")
	ErrCnpjNotFound           = errors.New("CNPJ não encontrado")
)

type BrasilApiService interface {
	GetStates() ([]resources.BrazilState, *merr.ResponseError)
	GetCnpjData(cnpj string) (*resources.BrasilApiCnpj, *merr.ResponseError)
	GetCepData(cep string) (*resources.BrasilApiCep, *merr.ResponseError)
}

type brasilApiService struct{}

func NewBrasilApiService() BrasilApiService {
	return &brasilApiService{}
}

func (s *brasilApiService) GetStates() ([]resources.BrazilState, *merr.ResponseError) {
	url := "https://brasilapi.com.br/api/ibge/uf/v1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchBrasilApi)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchBrasilApi)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchBrasilApi)
	}

	var states []resources.BrazilState
	if err := json.NewDecoder(resp.Body).Decode(&states); err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToParseBrasilApi)
	}

	return states, nil
}

func (s *brasilApiService) GetCnpjData(cnpj string) (*resources.BrasilApiCnpj, *merr.ResponseError) {
	cnpjDigits := onlyDigits(cnpj)

	if cnpjDigits == "" || len(cnpjDigits) != 14 {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidCnpj)
	}

	url := fmt.Sprintf("https://brasilapi.com.br/api/cnpj/v1/%s", cnpjDigits)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	var cnpjData resources.BrasilApiCnpj
	if err := json.NewDecoder(resp.Body).Decode(&cnpjData); err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	return &cnpjData, nil
}

func (s *brasilApiService) GetCepData(cep string) (*resources.BrasilApiCep, *merr.ResponseError) {
	cepDigits := onlyDigits(cep)

	if cepDigits == "" || len(cepDigits) != 8 {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidCepFormat)
	}

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cepDigits)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCepNotFound)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCepNotFound)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCepNotFound)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCepNotFound)
	}

	var cepData resources.BrasilApiCep
	if err := json.NewDecoder(resp.Body).Decode(&cepData); err != nil {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCepNotFound)
	}

	return &cepData, nil
}
