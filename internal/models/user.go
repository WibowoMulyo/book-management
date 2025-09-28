package models

import (
	"time"
)

type User struct {
	ID         int       `json:"id" db:"id"`
	Username   string    `json:"username" db:"username" validate:"required,min=3,max=50"`
	Password   string    `json:"password,omitempty" db:"password" validate:"required,min=6"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy string    `json:"modified_by" db:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      UserInfo  `json:"user"`
}

type UserInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
