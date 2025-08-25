package auth

type MeResource struct {
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type LoginResource struct {
	Token string     `json:"token"`
	User  MeResource `json:"user"`
}
