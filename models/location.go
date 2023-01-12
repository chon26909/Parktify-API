package models

import "github.com/google/uuid"

type Coordinate struct {
	Latitude  int64
	Longitude int64
}

type Location struct {
	Id     uuid.UUID `gorm:"primarKey;column:id"`
	Coords Coordinate
}
