package models

type SignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInSuccess struct {
	Account     Account `json:"account"`
	AccessToken string  `json:"access_token"`
}
