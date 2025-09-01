package auth

type MeResource struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	UserType string `json:"user_type" binding:"required,oneof=supplier buyer"`
}
