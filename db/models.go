package db

import (
	"app/pb"
	"time"
)

type MainComputer struct {
	ID        uint32    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	*pb.MainComputer
}

type BrakeManager struct {
	ID        uint32    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	*pb.BrakeManager
}

type Sensor struct {
	ID     uint32 `gorm:"primarykey" json:"id"`
	Name   string `json:"name"`
	Mesure string `json:"mesure"`
}

type SensorData struct {
	ID        uint32    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	*pb.SensorData
}
