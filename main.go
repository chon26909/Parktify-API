package main

import (
	"log"
	// "parktify/controllers"
	// "parktify/lib"
	// "parktify/middleware"
	// "parktify/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// // database
	// db := lib.NewMySqlConnection()
	// // cache
	// rdb := lib.NewRedisConnection()

	// userRepostory := repository.NewUserRepository(db, rdb)
	// userController := controllers.NewUserController(userRepostory)
	// authController := controllers.NewAuthController(userRepostory)

	// locationRepository := repository.NewLocationRepository(db, rdb)
	// locationController := controllers.NewLocationController(locationRepository)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// //auth
	// auth := app.Group("/auth")
	// auth.Post("/signin", authController.SignIn)
	// auth.Post("/signup", authController.SignUp)

	// //user
	// user := app.Group("/user")
	// user.Use(middleware.AuthorizationRequired())
	// user.Get("/", userController.GetAllUsers)
	// user.Post("/create", userController.CreateUser)
	// user.Put("/:id", userController.UpdateUser)
	// user.Delete("/:id", userController.DeleteUser)

	// //location
	// location := app.Group("/location")
	// location.Get("/", locationController.GetAllLocation)
	// location.Post("/create", locationController.CreateLocation)

	log.Fatal(app.Listen(":4000"))
}
