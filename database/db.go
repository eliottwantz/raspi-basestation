package database

import (
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

	db, err := sql.Open("sqlite", "./database/polyloop.db")
	if err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, err
	}

	return sqlc.New(db), nil
}
