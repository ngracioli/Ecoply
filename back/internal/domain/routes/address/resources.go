package address

import "ecoply/internal/domain/services"

type AddressResource struct {
	Cep           string `json:"cep"`
	Street        string `json:"street"`
	Neighborhood  string `json:"neighborhood"`
	City          string `json:"city"`
	State         string `json:"state"`
	StateInitials string `json:"state_initials"`
}

func addressDataToResource(data *services.AddressData) *AddressResource {
	return &AddressResource{
		Cep:           data.Cep,
		Street:        data.Street,
		Neighborhood:  data.Neighborhood,
		City:          data.City,
		State:         data.State,
		StateInitials: data.StateInitials,
	}
}
