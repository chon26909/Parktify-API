package main

import (
	"fmt"
	"log"
	"parktify/controllers"
	"parktify/lib"
	"parktify/middleware"
	"parktify/repository"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	initConfig()

	// database
	db := lib.NewMySqlConnection()
	// cache
	rdb := lib.NewRedisConnection()

	userRepostory := repository.NewUserRepository(db, rdb)
	userController := controllers.NewUserController(userRepostory)
	authController := controllers.NewAuthController(userRepostory)

	locationRepository := repository.NewLocationRepository(db, rdb)
	locationController := controllers.NewLocationController(locationRepository)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	//auth
	auth := app.Group("/auth")
	auth.Post("/signin", authController.SignIn)
	auth.Post("/signup", authController.SignUp)

	//user
	user := app.Group("/user")
	user.Use(middleware.AuthorizationRequired())
	user.Get("/", userController.GetAllUsers)
	user.Post("/create", userController.CreateUser)
	user.Put("/:id", userController.UpdateUser)
	user.Delete("/:id", userController.DeleteUser)

	//location
	location := app.Group("/location")
	location.Get("/", locationController.GetAllLocation)
	location.Post("/create", locationController.CreateLocation)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port"))))

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // อ่าน value จาก ENV variable
	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
