package repository

import (
	"parktify/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type LocationRepository interface {
	GetAllLocation() (locations []models.Location, err error)
	CreateLocation(location models.Location) error
	UpdateLocation(location models.Location) error
}

type locationRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewLocationRepository(db *gorm.DB, redis *redis.Client) LocationRepository {
	return &locationRepository{db, redis}
}

func (r *locationRepository) GetAllLocation() (locations []models.Location, err error) {

	err = r.db.Find(&locations).Error
	if err != nil {
		return nil, err
	}

	return locations, err
}

func (r *locationRepository) GetLocationById(id string) (location *models.Location, err error) {

	err = r.db.Where("id = ?", id).Find(&location).Error

	return location, err
}

func (r *locationRepository) CreateLocation(location models.Location) error {

	return r.db.Create(location).Error
}

func (r *locationRepository) UpdateLocation(location models.Location) error {
	return nil
}
