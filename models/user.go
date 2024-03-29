package models

import (
	"time"
)

// User represent the user table
type User struct {
	ID                   int64     `json:"id"                    db:"id"`
	Email                string    `json:"email"                 db:"email"`
	Password             string    `json:"password"              db:"password"`
	PasswordConfirmation string    `json:"password_confirmation"`
	Name                 string    `json:"name"                  db:"name"`
	AvatarURL            string    `json:"avatar_url"`
	Address              string    `json:"address"               db:"address"`
	Bio                  string    `json:"bio"                   db:"bio"`
	IsAdmin              bool      `json:"is_admin"              db:"is_admin"`
	CreatedAt            time.Time `json:"created_at"            db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"            db:"updated_at"`
}
