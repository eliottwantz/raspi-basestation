package database

import (
	"app/pb"
	"time"
)

//	type ControlData struct {
//		ID             int64     `json:"id"`
//		MainComputerID int64     `json:"main_computer_id"`
//		BrakeManagerID int64     `json:"brake_manager_id"`
//		Value          string    `json:"value"`
//		CreatedAt      time.Time `json:"created_at"`
//	}
type MainComputer struct {
	ID uint32 `gorm:"primarykey" json:"id"`
	*pb.MainComputer
}

type BrakeManager struct {
	ID uint32 `gorm:"primarykey" json:"id"`
	*pb.BrakeManager
}

// type SensorState struct {
// 	MainComputer
// 	BrakeManager
// }

// type SensorState struct {
// 	ID             uint `gorm:"primarykey"`
// 	CreatedAt      time.Time
// 	MainComputerID uint
// 	BrakeManagerID uint
// 	MainComputer   *MainComputer `gorm:"foreignKey:MainComputerID"`
// 	BrakeManager   *BrakeManager `gorm:" foreignKey:BrakeManagerID"`
// }

type Sensor struct {
	ID     uint32 `gorm:"primarykey" json:"id"`
	Name   string `json:"name"`
	Mesure string `json:"mesure"`
}

type SensorData struct {
	ID        uint32    `gorm:"primarykey" json:"id"`
	Value     float32   `json:"value"`
	SensorID  uint      `json:"sensor_id"`
	CreatedAt time.Time `json:"created_at"`
}
