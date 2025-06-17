package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PasswordHash string    `json:"-"` // Password hash is not included in JSON output
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
