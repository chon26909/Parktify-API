package repository

import (
	"context"
	"parktify/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
}

type userRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{db, redis}
}

func (r *userRepository) CreateUser(user *models.User) error {

	r.db.Create(user)
	return nil
}

func (r *userRepository) GetAllUsers() (users []*models.User, err error) {

	exist := r.redis.Get(context.TODO(), "A")

	if exist == nil {
		err := r.db.Table("users").Find(&users).Error
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
