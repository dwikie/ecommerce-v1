package entity

import (
	"database/sql"
	"time"
)

// Entity
type Account struct {
	Id        string    `gorm:"primaryKey;not null"`
	Fullname  string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Phone     *string
	Password  string        `gorm:"not null"`
	CreatedAt time.Time     `gorm:"not null"`
	UpdatedAt *sql.NullTime `gorm:"index"`
}
