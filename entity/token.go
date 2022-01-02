package entity

import (
	"time"

	"gorm.io/gorm"
)

// Entity
type Token struct {
	Id        uint64
	Issuer    string
	UserAgent string
	Token     string
	Type      uint8
	Expire    time.Time
	AccountId string
	gorm.Model
}
