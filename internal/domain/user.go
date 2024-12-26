package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string     `json:"name"`
	Password        string     `json:"-"`
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	EmailVerifiedAt *time.Time `json:"phoneVerifiedAt"`
}
