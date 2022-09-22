package api

import (
	"app/db"
	"app/pb"

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
	ss := pb.SensorState{
		MainComputer: mc.MainComputer,
		BrakeManager: bm.BrakeManager,
	}
	return c.JSON(&ss)
}
