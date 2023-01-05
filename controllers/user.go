package controllers

import (
	"fmt"
	"parktify/dto"
	"parktify/repository"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
}

type userController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{userRepository: userRepo}
}

func (r *userController) GetAllUsers(ctx *fiber.Ctx) error {

	fmt.Println("get all users")

	users, _ := r.userRepository.GetUsers()

	responseUsers := []*dto.User{}
	for _, user := range users {
		responseUsers = append(responseUsers, &dto.User{
			UserID:       user.UserID,
			Username:     "",
			FirstName:    "",
			LastName:     "",
			MobileNumber: "",
			Email:        user.Email,
		})
	}

	fmt.Println("users ", users)

	return ctx.JSON(fiber.Map{"message": "ok", "data": responseUsers})
}
