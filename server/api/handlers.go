package api

import (
	"app/db"

	"github.com/gofiber/fiber/v2"
)

func HandleSensorState(c *fiber.Ctx) error {
	mc, err := db.GetLatestMainComputer()
	if err != nil {
		return fiber.ErrNotFound
	}
	bm, err := db.GetLatestBrakeManager()
	if err != nil {
		return fiber.ErrNotFound
	}
	ss := db.SensorState{
		MainComputer: mc,
		BrakeManager: bm,
	}
	return c.JSON(&ss)
}

func HandleMainComputer(c *fiber.Ctx) error {
	mc, err := db.GetLatestMainComputer()
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&mc)
}

func HandleBrakeManager(c *fiber.Ctx) error {
	bm, err := db.GetLatestBrakeManager()
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&bm)
}
