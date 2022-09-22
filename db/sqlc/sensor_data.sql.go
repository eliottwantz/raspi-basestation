// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: sensor_data.sql

package sqlc

import (
	"context"
	"time"
)

const createSensorData = `-- name: CreateSensorData :one
INSERT INTO sensor_datas (created_at, sensor_id, value)
VALUES (?, ?, ?)
RETURNING id, created_at, sensor_id, value
`

type CreateSensorDataParams struct {
	CreatedAt time.Time `json:"created_at"`
	SensorID  int64     `json:"sensor_id"`
	Value     float64   `json:"value"`
}

func (q *Queries) CreateSensorData(ctx context.Context, arg CreateSensorDataParams) (SensorData, error) {
	row := q.db.QueryRowContext(ctx, createSensorData, arg.CreatedAt, arg.SensorID, arg.Value)
	var i SensorData
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.SensorID,
		&i.Value,
	)
	return i, err
}

const deleteSensorData = `-- name: DeleteSensorData :exec
DELETE FROM sensor_datas
WHERE id = ?
`

func (q *Queries) DeleteSensorData(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteSensorData, id)
	return err
}

const getAllSensorData = `-- name: GetAllSensorData :many
SELECT id, created_at, sensor_id, value
FROM sensor_datas
ORDER BY id
`

func (q *Queries) GetAllSensorData(ctx context.Context) ([]SensorData, error) {
	rows, err := q.db.QueryContext(ctx, getAllSensorData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SensorData
	for rows.Next() {
		var i SensorData
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.SensorID,
			&i.Value,
		); err != nil {
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

const getLatestSensorData = `-- name: GetLatestSensorData :one
SELECT id, created_at, sensor_id, value
FROM sensor_datas
WHERE sensor_id = ?
ORDER BY id DESC
LIMIT 1
`

func (q *Queries) GetLatestSensorData(ctx context.Context, sensorID int64) (SensorData, error) {
	row := q.db.QueryRowContext(ctx, getLatestSensorData, sensorID)
	var i SensorData
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.SensorID,
		&i.Value,
	)
	return i, err
}
