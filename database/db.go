package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/polyloop.sqlite3"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Exec("PRAGMA journal_mode=off").Error
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&MainComputer{}, &BrakeManager{}, &Sensor{}, &SensorData{})
	if err != nil {
		panic(err)
	}
	err = db.Create(&Sensor{Name: "premierSensor", Mesure: "km/h"}).Error
	if err != nil {
		panic(err)
	}
	return db
}
