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
)

type LocationController interface {
	GetAllLocation(ctx *fiber.Ctx) error
	CreateLocation(ctx *fiber.Ctx) error
}

type locationController struct {
	locationRepository repository.LocationRepository
}

func NewLocationController(locationRepo repository.LocationRepository) LocationController {
	return &locationController{locationRepository: locationRepo}
}

func (r *locationController) GetAllLocation(ctx *fiber.Ctx) error {

	locations, err := r.locationRepository.GetAllLocation()

	locationResponse := []*dto.LocationResponse{}

	for _, item := range locations {
		locationResponse = append(locationResponse, &dto.LocationResponse{
			Latitude:    item.Latitude,
			Longitude:   item.Longitude,
			Title:       item.Title,
			Description: item.Description,
		})
	}

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"message": "ok", "data": locationResponse})
}

func (r *locationController) CreateLocation(ctx *fiber.Ctx) error {
	var body dto.CreateLocationRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return err
	}

	res, _ := utils.UploadImage(body.Image)

	fmt.Print("response ", res)
	return nil

	newLocation := models.Location{
		LocationID:  uuid.New(),
		Latitude:    body.Latitude,
		Longitude:   body.Longitude,
		Title:       body.Title,
		Description: body.Description,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	err = r.locationRepository.CreateLocation(newLocation)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "created", "data": body})
}
