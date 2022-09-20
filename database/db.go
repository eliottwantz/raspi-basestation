package database

import (
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() *gorm.DB {
	db, err := gorm.Open(gormsqlite.Open("./database/polyloop.sqlite3"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.Exec("PRAGMA journal_mode=off").Error
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&MainComputer{}, &BrakeManager{})
	if err != nil {
		panic(err)
	}
	return db
}
