package auth

type MeResource struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

type LoginResource struct {
	Token string     `json:"token"`
	User  MeResource `json:"user"`
}
