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
	Cnpj            string          `json:"cnpj" binding:"required,len=14,numeric,cnpj"`
	UserType        string          `json:"user_type" binding:"required,oneof=supplier buyer"`
	Address         *AddressRequest `json:"address" binding:"required"`
}

type AddressRequest struct {
	Cep     string `json:"cep" binding:"required,len=8,numeric"`
	State   string `json:"state" binding:"required,min=2,max=50"`
	City    string `json:"city" binding:"required,min=2,max=100"`
	Country string `json:"country" binding:"required,min=2,max=50"`
}

type AvailabilityRequest struct {
	Type  string `form:"type" json:"type" binding:"required,oneof=email cnpj"`
	Value string `form:"value" json:"value" binding:"required"`
}
