package main

import (
	"log"
	"parktify/controllers"
	"parktify/lib"
	"parktify/repository"

	// "parktify/lib"
	// "parktify/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := lib.NewMySqlConnection()

	_ = db

	rdb := lib.NewRedisConnection()

	userRepostory := repository.NewUserRepository(db, rdb)
	userController := controllers.NewUserController(userRepostory)

	_ = userRepostory

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	api := app.Group("/api")

	//user
	api.Get("/user", userController.GetAllUsers)

	log.Fatal(app.Listen(":4000"))
}
