package dto

import "github.com/google/uuid"

type User struct {
	UserID       uuid.UUID `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	FirstName    string    `json:"firstname,omitempty"`
	LastName     string    `json:"lastname,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	Email        string    `json:"email,omitempty"`
}

type RequestUser struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
}
