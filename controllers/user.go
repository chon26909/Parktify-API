package controllers

import (
	"fmt"
	"parktify/dto"
	"parktify/models"
	"parktify/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type userController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) UserController {
	return &userController{userRepository: userRepo}
}

func (r *userController) GetAllUsers(ctx *fiber.Ctx) error {

	fmt.Println("get all users")

	users, err := r.userRepository.GetUsers()
	if err != nil {
		return err
	}

	responseUsers := []*dto.User{}
	for _, user := range users {
		responseUsers = append(responseUsers, &dto.User{
			UserID:       user.UserID,
			Username:     user.Username,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			MobileNumber: "",
			Email:        user.Email,
		})
	}

	return ctx.JSON(fiber.Map{"message": "ok", "data": responseUsers})
}

func (r *userController) CreateUser(ctx *fiber.Ctx) error {

	var body dto.RequestUser
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	user := models.User{
		UserID:       uuid.New(),
		Username:     "",
		FirstName:    "",
		LastName:     "",
		MobileNumber: "",
		Email:        body.Email,
		Password:     "",
		Created:      time.Now(),
		Updated:      time.Now(),
	}

	fmt.Println("user body", user)

	r.userRepository.CreateUser(user)

	return ctx.JSON(fiber.Map{"message": "ok", "data": user})
}

func (r *userController) UpdateUser(ctx *fiber.Ctx) error {

	uid := uuid.MustParse(ctx.Params("id"))

	var body dto.RequestUser
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	user := models.User{
		UserID:    uid,
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	r.userRepository.UpdateUser(user)

	return ctx.JSON(fiber.Map{"message": "updated", "data": user})
}

func (r *userController) DeleteUser(ctx *fiber.Ctx) error {

	uid := uuid.MustParse(ctx.Params("id"))

	r.userRepository.DeleteUser(uid)

	return nil
}
