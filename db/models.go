package db

import "app/db/sqlc"

type SensorState struct {
	BrakeManager *sqlc.BrakeManager `json:"brakeManager"`
	MainComputer *sqlc.MainComputer `json:"mainComputer"`
}
