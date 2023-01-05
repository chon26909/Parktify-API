package main

import (
	"log"
	// "parktify/lib"
	// "parktify/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// db := lib.NewMySqlConnection()
	// rdb := lib.NewRedisConnection()

	// userRepostory := repository.NewUserRepository(db, rdb)

	// _ = userRepostory

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Fatal(app.Listen(":4000"))
}
