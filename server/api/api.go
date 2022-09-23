package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"

	json "github.com/goccy/go-json"
)

func Start() {
	app := fiber.New(fiber.Config{
		Prefork:     false,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New(), etag.New(), logger.New())
	RegisterRoutes(app)
	log.Fatal(app.Listen(":8000"))
}

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/SensorState", HandleSensorState)
	api.Get("/BrakeManager", HandleBrakeManager)
	api.Get("/MainComputer", HandleMainComputer)

	api.Get("/SensorData/:id", HandleSensorData)

	app.Get("/ws", HandleWS())
}
