package api

import (
	"app/db"

	"github.com/gofiber/fiber/v2"
)

func HandleSensorState(c *fiber.Ctx) error {
	mc, err := db.Queries.GetLatestMainComputer(db.Ctx)
	if err != nil {
		return fiber.ErrNotFound
	}
	bm, err := db.Queries.GetLatestBrakeManager(db.Ctx)
	if err != nil {
		return fiber.ErrNotFound
	}
	ss := db.SensorState{
		MainComputer: &mc,
		BrakeManager: &bm,
	}
	return c.JSON(&ss)
}

func HandleMainComputer(c *fiber.Ctx) error {
	mc, err := db.Queries.GetLatestMainComputer(db.Ctx)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&mc)
}

func HandleBrakeManager(c *fiber.Ctx) error {
	bm, err := db.Queries.GetLatestBrakeManager(db.Ctx)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&bm)
}

func HandleSensorData(c *fiber.Ctx) error {
	sID, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	sd, err := db.Queries.GetLatestSensorData(db.Ctx, int64(sID))
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&sd)
}