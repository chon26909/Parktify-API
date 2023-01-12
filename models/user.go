package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"primaryKey;column:id"`
	Username     string    `gorm:"column:username"`
	FirstName    string    `gorm:"column:firstname"`
	LastName     string    `gorm:"column:lastname"`
	MobileNumber string    `gorm:"column:mobile_number"`
	Email        string
	Password     string
	Created      time.Time
	Updated      time.Time `gorm:"autoUpdateTime"`
}
