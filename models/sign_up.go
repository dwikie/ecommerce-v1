package models

import "time"

type SignUp struct {
	Fullname  string    `json:"fullname" validate:"required"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
}
