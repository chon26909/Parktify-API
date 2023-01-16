package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	LocationID  uuid.UUID `gorm:"primarKey;column:id"`
	Latitude    float64
	Longitude   float64
	Title       string
	Description string
	CreateBy    uuid.UUID
	Created     time.Time
	Updated     time.Time `gorm:"autoUpdateTime"`
}
