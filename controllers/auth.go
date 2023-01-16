package controllers

import (
	"fmt"
	"parktify/dto"
	"parktify/models"
	"parktify/repository"
	"parktify/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthController interface {
	SignUp(ctx *fiber.Ctx) error
	SignIn(cys *fiber.Ctx) error
}

type authController struct {
	userRepository repository.UserRepository
}

func NewAuthController(userRepo repository.UserRepository) AuthController {
	return &authController{userRepository: userRepo}
}

func (r *authController) SignUp(ctx *fiber.Ctx) error {

	var body *dto.SignUpRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	_, err = r.userRepository.GetUserByEmail(body.Email)
	if err == gorm.ErrRecordNotFound {

		hashPassword, _ := utils.EncryptPassword(body.Password)

		newUser := models.User{
			UserID:       uuid.New(),
			Username:     body.Username,
			FirstName:    body.Firstname,
			LastName:     body.Lastname,
			MobileNumber: "",
			Email:        body.Email,
			Password:     hashPassword,
			Created:      time.Now(),
			Updated:      time.Now(),
		}

		err := r.userRepository.CreateUser(newUser)
		if err != nil {
			return err
		}

		token, err := utils.GenerateTokenAuth(newUser.UserID)
		if err != nil {
			fmt.Println("token error:", err)
		}

		return ctx.Status(201).JSON(fiber.Map{
			"message": "created",
			"token":   token,
		})

	}

	return ctx.Status(400).JSON(fiber.Map{
		"message": "Invalid username or password",
	})
}

func (r *authController) SignIn(ctx *fiber.Ctx) error {

	var body *dto.SignInRequest
	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	user, err := r.userRepository.GetUserByEmail(body.Email)
	if err != nil {
		return err
	}

	token, err := utils.GenerateTokenAuth(user.UserID)
	if err != nil {
		fmt.Println("token error:", err)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}
