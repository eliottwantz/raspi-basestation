package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database/polyloop.db"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&MainComputer{}, &BrakeManager{})
	if err != nil {
		panic(err)
	}
	return db, nil
}

func GetCount(db *gorm.DB) int64 {
	var count int64
	db.Model(&MainComputer{}).Count(&count)
	return count
}
