package database

import (
	"app/internal"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Open() {
	db, err := gorm.Open(sqlite.Open("./database/polyloop.sqlite3"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	internal.HandleError(err)
	internal.HandleError(db.Exec("PRAGMA journal_mode=off").Error)
	internal.HandleError(db.AutoMigrate(&MainComputer{}, &BrakeManager{}, &Sensor{}, &SensorData{}))
	internal.HandleError(db.Create(&Sensor{Name: "premierSensor", Mesure: "km/h"}).Error)
	DB = db
}
