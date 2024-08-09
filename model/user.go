package model

import (
	"time"
)

type User struct {
	GormGeneric

	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"-"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
}
