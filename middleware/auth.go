package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte("parktify"),
		SigningMethod: "HS256",
		TokenLookup:   "header:Authorization",
		SuccessHandler: func(ctx *fiber.Ctx) error {
			ctx.Next()
			return nil
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
			return nil
		},
	})
}
