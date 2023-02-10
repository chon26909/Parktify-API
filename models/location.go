package models

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	LocationID  uuid.UUID `gorm:"primarKey;column:id;type:uuid;default:uuid_generate_v4()"`
	Latitude    float64
	Longitude   float64
	Title       string
	Description string
	TimeOpem    time.Time
	TimeClose   time.Time
	CreateBy    uuid.UUID
	Created     time.Time
	Updated     time.Time `gorm:"autoUpdateTime"`
}
