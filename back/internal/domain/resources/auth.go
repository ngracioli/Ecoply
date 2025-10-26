package resources

type Me struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	UserType string  `json:"user_type"`
	Address  Address `json:"address"`
	Agent    Agent   `json:"agent"`
}

type Address struct {
	Cep          string `json:"cep"`
	StateInitial string `json:"state_initial"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
}

type Agent struct {
	Cnpj          string `json:"cnpj"`
	CceeCode      string `json:"ccee_code"`
	SubmarketName string `json:"submarket_name"`
	CompanyName   string `json:"company_name"`
}

type Login struct {
	Token string `json:"token"`
	User  Me     `json:"user"`
}
