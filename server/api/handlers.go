package api

import (
	"app/db"
	"app/db/sqlc"
	"app/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

const (
	SENSOR_STATE string = "state"
	SENSOR_DATA  string = "data"
)

type SensorStateWSMessage struct {
	Type string `json:"type"`
	*db.SensorState
}
type SensorDataWSMessage struct {
	Type       string           `json:"type"`
	SensorData *sqlc.SensorData `json:"sensorData"`
}

func HandleWS() func(*fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		for {
			select {
			case ss := <-server.Wsss:
				// c.WriteJSON(&SensorStateWSMessage{
				// 	Type:        SENSOR_STATE,
				// 	SensorState: ss,
				// })
				c.WriteJSON(&ss)
			case sd := <-server.Wssd:
				// c.WriteJSON(&SensorDataWSMessage{
				// 	Type:       SENSOR_DATA,
				// 	SensorData: sd,
				// })
				c.WriteJSON(&sd)
			}
		}
	})
}
