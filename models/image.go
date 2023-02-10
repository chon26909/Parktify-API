package models

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	LocationID uuid.UUID `gorm:"column:id"`
	Image      string
	Updated    time.Time `gorm:"autoUpdateTime"`
}
