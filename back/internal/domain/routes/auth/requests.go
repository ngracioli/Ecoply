package auth

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignUpRequest struct {
	Name            string          `json:"name" binding:"required,min=2,max=100"`
	Email           string          `json:"email" binding:"required,email"`
	Password        string          `json:"password" binding:"required,min=8,max=50,containsuppercase,containslowercase,containsdigit,containsspecial"`
	ConfirmPassword string          `json:"confirm_password" binding:"required,eqfield=Password"`
	UserType        string          `json:"user_type" binding:"required,oneof=supplier buyer"`
	Agent           *AgentRequest   `json:"agent" binding:"required"`
	Address         *AddressRequest `json:"address" binding:"required"`
}

type AgentRequest struct {
	Cnpj          string `json:"cnpj" binding:"required,len=14,numeric,cnpj"`
	CompanyName   string `json:"company_name" binding:"required,min=2,max=255"`
	CceeCode      string `json:"ccee_code" binding:"required,min=1,max=50"`
	SubmarketName string `json:"submarket_name" binding:"required,oneof=N S SE_CO NE"`
}

type AddressRequest struct {
	Cep           string `json:"cep" binding:"required,len=8,numeric"`
	Number        string `json:"number" binding:"required,min=1,max=20"`
	Complement    string `json:"complement" binding:"max=50"`
	Street        string `json:"street" binding:"required,min=2,max=255"`
	Neighborhood  string `json:"neighborhood" binding:"required,min=2,max=255"`
	City          string `json:"city" binding:"required,min=2,max=255"`
	State         string `json:"state" binding:"required,min=2,max=255"`
	StateInitials string `json:"state_initials" binding:"required,len=2,uppercase"`
}

type AvailabilityRequest struct {
	Type  string `form:"type" json:"type" binding:"required,oneof=email cnpj ccee"`
	Value string `form:"value" json:"value" binding:"required"`
}
