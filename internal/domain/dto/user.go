package dto

type CreateUserInput struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

type CreateUserOutput struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

type AuthenticateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateUserOutput struct {
	AccessToken string `json:"access_token"`
}
