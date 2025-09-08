package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/mlog"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrFailedToFetchAddress = errors.New("failed to fetch address from external service")
	ErrAddressNotFound      = errors.New("address not found")
	ErrInvalidCep           = errors.New("invalid CEP")
	ErrFailedToParseAddress = errors.New("failed to parse address response")
)

type AddressData struct {
	Error         string `json:"erro,omitempty"`
	Cep           string `json:"cep"`
	State         string `json:"estado"`
	City          string `json:"localidade"`
	Neighborhood  string `json:"bairro"`
	Street        string `json:"logradouro"`
	StateInitials string `json:"uf"`
}

func LoadAddressByCep(cep string) (*AddressData, *merr.ResponseError) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchAddress)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrAddressNotFound)
	}

	var addressData AddressData
	if err := json.NewDecoder(resp.Body).Decode(&addressData); err != nil {
		mlog.Log("Failed to decode address data: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToParseAddress)
	}

	// Eu odeio o viacep com todas as minhas forças, morra viacep
	if addressData.Error == "true" {
		return nil, merr.NewResponseError(http.StatusBadRequest, ErrInvalidCep)
	}

	return &addressData, nil
}
