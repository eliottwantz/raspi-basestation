-- name: GetLatestSensor :one
SELECT *
FROM sensors
ORDER BY id DESC
LIMIT 1;
-- name: GetAllSensor :many
SELECT *
FROM sensors
ORDER BY id;
-- name: CreateSensor :one
INSERT INTO sensors (name, mesure)
VALUES (?, ?)
RETURNING *;
-- name: DeleteSensor :exec
DELETE FROM sensors
WHERE id = ?;