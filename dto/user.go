package dto

import "github.com/google/uuid"

type User struct {
	UserID       uuid.UUID `json:"user_id,omitempty"`
	Username     string    `json:"username,omitempty"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	Email        string    `json:"email,omitempty"`
}
