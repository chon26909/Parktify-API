package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {

	salt := bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)

	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateTokenAuth(uid uuid.UUID) (string, error) {

	fmt.Println("uid", uid)

	claim := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte("parktify"))
}

func GetIdFromToken(ctx *fiber.Ctx) (uuid.UUID, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	uid := claims["id"].(string)

	return uuid.MustParse(uid), nil
}
