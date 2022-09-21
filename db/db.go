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
	internal.HandleError(err)
	internal.HandleError(DB.Exec("PRAGMA journal_mode=off").Error)
	internal.HandleError(DB.AutoMigrate(&MainComputer{}, &BrakeManager{}, &Sensor{}, &SensorData{}))
	internal.HandleError(DB.Create(&Sensor{Name: "premierSensor", Mesure: "km/h"}).Error)
}
