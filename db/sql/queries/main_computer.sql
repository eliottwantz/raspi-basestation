-- name: GetLatestMainComputer :one
SELECT *
FROM main_computers
ORDER BY id DESC
LIMIT 1;
-- name: GetAllMainComputer :many
SELECT *
FROM main_computers
ORDER BY id;
-- name: CreateMainComputer :one
INSERT INTO main_computers (created_at, state)
VALUES (?, ?)
RETURNING *;
-- name: DeleteMainComputer :exec
DELETE FROM main_computers
WHERE id = ?;