package services

import (
	"context"
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/resources"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type CceeService interface {
	GetAgentsByCnpj(cnpj string) ([]*resources.CceeAgent, *merr.ResponseError)
}

type cceeService struct{}

func NewCceeService() CceeService {
	return &cceeService{}
}

type cceeResult struct {
	agents []*resources.CceeAgent
	err    *merr.ResponseError
}

func (s *cceeService) GetAgentsByCnpj(cnpj string) ([]*resources.CceeAgent, *merr.ResponseError) {
	resultChan := make(chan cceeResult, 1)

	go func() {
		agents, err := s.fetchAgentsByCnpj(cnpj)
		resultChan <- cceeResult{agents: agents, err: err}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))
	defer cancel()

	select {
	case result := <-resultChan:
		return result.agents, result.err
	case <-ctx.Done():
		return nil, merr.NewResponseError(http.StatusRequestTimeout, ErrFailedToFetchCcee)
	}
}

func (s *cceeService) fetchAgentsByCnpj(cnpj string) ([]*resources.CceeAgent, *merr.ResponseError) {
	cnpjDigits := onlyDigits(cnpj)

	if cnpjDigits == "" || len(cnpjDigits) != 14 {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidCnpj)
	}

	url := fmt.Sprintf(
		"https://dadosabertos.ccee.org.br/api/3/action/datastore_search?resource_id=71169d34-7171-47bb-8217-4ff140fed41d&q=%s",
		cnpjDigits,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchCcee)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchCcee)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchCcee)
	}

	var apiResponse resources.CceeApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToParseCcee)
	}

	if !apiResponse.Success {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrFailedToFetchCcee)
	}

	var validRecords []resources.CceeRecord
	for _, record := range apiResponse.Result.Records {
		if onlyDigits(record.Cnpj) == cnpjDigits {
			validRecords = append(validRecords, record)
		}
	}

	if len(validRecords) == 0 {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrCceeNotFound)
	}

	agents := make([]*resources.CceeAgent, 0, len(validRecords))
	for _, record := range validRecords {
		agents = append(agents, &resources.CceeAgent{
			Label:     fmt.Sprintf("%d - %s", record.CodPerfAgente, record.SiglaPerfAgente),
			Value:     fmt.Sprintf("%d", record.CodPerfAgente),
			Submarket: mapSubmarketToCode(record.Submercado),
		})
	}

	return agents, nil
}

func onlyDigits(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}

func mapSubmarketToCode(submarket string) string {
	submarketMap := map[string]string{
		"SUDESTE":  "SE_CO",
		"SUL":      "S",
		"NORDESTE": "NE",
		"NORTE":    "N",
	}

	if code, ok := submarketMap[submarket]; ok {
		return code
	}

	return submarket
}
