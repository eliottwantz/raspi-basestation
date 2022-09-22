package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	api.Get("/sensorstate", HandleSensorState)
}
