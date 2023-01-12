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

	// database
	db := lib.NewMySqlConnection()
	// cache
	rdb := lib.NewRedisConnection()

	userRepostory := repository.NewUserRepository(db, rdb)
	userController := controllers.NewUserController(userRepostory)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	// api := app.Group("/api")

	//user
	app.Get("/user", userController.GetAllUsers)
	app.Post("/user/create", userController.CreateUser)
	app.Put("/user/:id", userController.UpdateUser)
	app.Delete("/user/:id", userController.DeleteUser)

	log.Fatal(app.Listen(":4000"))
}
