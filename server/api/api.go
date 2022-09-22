package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	api.Get("/SensorState", HandleSensorState)
	api.Get("/BrakeManager", HandleBrakeManager)
	api.Get("/MainComputer", HandleMainComputer)

	api.Get("/SensorData/:id", HandleSensorData)
}
