package repository

import (
	"parktify/models"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() (users []models.User, err error)
	UpdateUser(user models.User) error
	DeleteUser(uid uuid.UUID) error
}

type userRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{db, redis}
}

func (r *userRepository) CreateUser(user models.User) error {

	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {

	return nil, nil
}

func (r *userRepository) GetUsers() (users []models.User, err error) {

	err = r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.db.Model(&user).Updates(user).Error
}

func (r *userRepository) DeleteUser(uid uuid.UUID) error {
	return r.db.Delete(&models.User{}, uid).Error
}
