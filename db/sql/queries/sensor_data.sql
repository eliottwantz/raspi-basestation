-- name: GetLatestSensorData :one
SELECT *
FROM sensor_datas
WHERE sensor_id = ?
ORDER BY id DESC
LIMIT 1;
-- name: GetAllSensorData :many
SELECT *
FROM sensor_datas
ORDER BY id;
-- name: CreateSensorData :one
INSERT INTO sensor_datas (created_at, sensor_id, value)
VALUES (?, ?, ?)
RETURNING *;
-- name: DeleteSensorData :exec
DELETE FROM sensor_datas
WHERE id = ?;