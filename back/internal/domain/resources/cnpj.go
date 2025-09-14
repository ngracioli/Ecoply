package resources

type Company struct {
	Cnpj    string         `json:"cnpj"`
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
