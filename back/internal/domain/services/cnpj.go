package services

import (
	"ecoply/internal/domain/merr"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrFailedToFetchCnpj = errors.New("failed to fetch cnpj from external service")
	ErrCnpjNotFound      = errors.New("cnpj not found")
	ErrInvalidCep        = errors.New("invalid CEP")
	ErrFailedToParseCnpj = errors.New("failed to parse cnpj response")
)

type CnpjData struct {
	TaxId   string      `json:"taxId"`
	Company CnpjCompany `json:"company"`
	Address CnpjAddress `json:"address"`
}

type CnpjCompany struct {
	Name string `json:"name"`
}

type CnpjAddress struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
	Details  string `json:"details"`
}

func LoadCnpjData(cnpj string) (*CnpjData, *merr.ResponseError) {
	resp, err := http.Get("https://open.cnpja.com/office/" + cnpj)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchCnpj)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	var cnpjData CnpjData
	if err := json.NewDecoder(resp.Body).Decode(&cnpjData); err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToParseCnpj)
	}

	if cnpjData.TaxId == "" || cnpjData.Company.Name == "" {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCnpjNotFound)
	}

	return &cnpjData, nil
}
