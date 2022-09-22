package db

import (
	"app/db/sqlc"
	"app/internal"
	"context"
	"database/sql"

	_ "embed"

	_ "modernc.org/sqlite"
)

var (
	//go:embed sql/schema.sql
	schema  string
	DB      *sql.DB
	Queries *sqlc.Queries
	Ctx     context.Context
)

func Open() {
	Ctx = context.Background()
	var err error
	DB, err = sql.Open("sqlite", "polyloop.sqlite3")
	internal.FatalError(err)
	_, err = DB.ExecContext(Ctx, "PRAGMA journal_mode=off")
	internal.FatalError(err)
	// create tables
	_, err = DB.ExecContext(Ctx, schema)
	internal.FatalError(err)
	Queries = sqlc.New(DB)
	_, err = Queries.CreateSensor(Ctx, sqlc.CreateSensorParams{
		Name:   "premierSensor",
		Mesure: "km/h",
	})
	internal.NotFatalError(err)
}
