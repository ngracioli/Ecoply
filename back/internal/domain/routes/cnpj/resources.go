package address

import "ecoply/internal/domain/services"

type CompanyResource struct {
	TaxId   string         `json:"taxId"`
	Name    string         `json:"name"`
	Address CompanyAddress `json:"address"`
}

type CompanyAddress struct {
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Complement   string `json:"complement"`
}

func companyDataToResource(data *services.CnpjData) *CompanyResource {
	return &CompanyResource{
		TaxId: data.TaxId,
		Name:  data.Company.Name,
		Address: CompanyAddress{
			Cep:          data.Address.Zip,
			Street:       data.Address.Street,
			Number:       data.Address.Number,
			Neighborhood: data.Address.District,
			City:         data.Address.City,
			State:        data.Address.State,
			Complement:   data.Address.Details,
		},
	}
}
