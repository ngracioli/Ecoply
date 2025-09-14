package resources

type Me struct {
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type Login struct {
	Token string `json:"token"`
	User  Me     `json:"user"`
}
