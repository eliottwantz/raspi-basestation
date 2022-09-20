// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: main_computer.sql

package sqlc

import (
	"context"
)

const createMainComputer = `-- name: CreateMainComputer :one
INSERT INTO
  main_computers (state)
VALUES
  (?) RETURNING id, state
`

func (q *Queries) CreateMainComputer(ctx context.Context, state int64) (MainComputer, error) {
	row := q.db.QueryRowContext(ctx, createMainComputer, state)
	var i MainComputer
	err := row.Scan(&i.ID, &i.State)
	return i, err
}

const deleteMainComputer = `-- name: DeleteMainComputer :exec
DELETE FROM
  main_computers
WHERE
  id = ?
`

func (q *Queries) DeleteMainComputer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMainComputer, id)
	return err
}

const getMainComputer = `-- name: GetMainComputer :one
SELECT
  id, state
FROM
  main_computers
WHERE
  id = ?
LIMIT
  1
`

func (q *Queries) GetMainComputer(ctx context.Context, id int64) (MainComputer, error) {
	row := q.db.QueryRowContext(ctx, getMainComputer, id)
	var i MainComputer
	err := row.Scan(&i.ID, &i.State)
	return i, err
}

const getMainComputerCount = `-- name: GetMainComputerCount :one
SELECT
  COUNT(*)
FROM
  main_computers
`

func (q *Queries) GetMainComputerCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getMainComputerCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMainComputers = `-- name: GetMainComputers :many
SELECT
  id, state
FROM
  main_computers
ORDER BY
  id
`

func (q *Queries) GetMainComputers(ctx context.Context) ([]MainComputer, error) {
	rows, err := q.db.QueryContext(ctx, getMainComputers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MainComputer
	for rows.Next() {
		var i MainComputer
		if err := rows.Scan(&i.ID, &i.State); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
