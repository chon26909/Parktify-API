package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"primaryKey;column:id"`
	Username     string    `gorm:"column:username"`
	FirstName    string    `gorm:"column:firstname"`
	LastName     string    `gorm:"column:lastname"`
	MobileNumber string    `gorm:"column:mobile_number"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Created      int64     `gorm:"autoCreateTime"`
	Updated      int64     `gorm:"autoUpdateTime"`
}
