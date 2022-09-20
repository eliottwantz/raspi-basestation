package database

import (
	// gormsqlite "gorm.io/driver/sqlite"
	// "gorm.io/gorm"

	"app/database/sqlc"
	"context"
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

var (
	//go:embed sql/schema.sql
	schema string
)

func Open() (*sqlc.Queries, error) {
	ctx := context.Background()

	// db, err := sql.Open("sqlite", ":memory:") // For tests purpuses
	db, err := sql.Open("sqlite", "./database/polyloop.sqlite3")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	return sqlc.New(db), nil
}

// func Migrate() {
// 	db, err := gorm.Open(gormsqlite.Open("./database/polyloop.db"), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 		PrepareStmt:            true,
// 	})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	err = db.AutoMigrate(&MainComputer{}, &BrakeManager{})
// 	if err != nil {
// 		panic(err)
// 	}
// }
