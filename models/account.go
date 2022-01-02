package models

import (
	"time"
)

type Account struct {
	Id        string     `json:"id"`
	Fullname  string     `json:"fullname"`
	Birthdate time.Time  `json:"birthdate"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
