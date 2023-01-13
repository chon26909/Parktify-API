package controllers

import (
	"fmt"
	"parktify/dto"
	"parktify/repository"

	"github.com/gofiber/fiber/v2"
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

	return nil
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
	fmt.Print("user", user)
	return nil
}
