package db

import (
	"app/internal"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Open() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./db/polyloop.sqlite3"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	internal.FatalError(err)
	internal.FatalError(DB.Exec("PRAGMA journal_mode=off").Error)
	internal.FatalError(DB.AutoMigrate(&MainComputer{}, &BrakeManager{}, &Sensor{}, &SensorData{}))
	internal.FatalError(DB.Create(&Sensor{Name: "premierSensor", Mesure: "km/h"}).Error)
}

func GetLatestMainComputer() (*MainComputer, error) {
	var mc MainComputer
	err := DB.Last(&mc).Error
	return &mc, err
}

func GetLatestBrakeManager() (*BrakeManager, error) {
	var bm BrakeManager
	err := DB.Last(&bm).Error
	return &bm, err
}

func GetLatestSensorData() (*SensorData, error) {
	var sd SensorData
	err := DB.Last(&sd).Error
	return &sd, err
}
