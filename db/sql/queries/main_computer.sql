-- name: GetMainComputerState :one
SELECT
  *
FROM
  main_computer
WHERE
  id = ?
LIMIT
  1;

-- name: GetMainComputerStates :many
SELECT
  *
FROM
  main_computer
ORDER BY
  id;

-- name: CreateMainComputerState :one
INSERT INTO
  main_computer (state)
VALUES
  (?) RETURNING *;

-- name: DeleteMainComputerState :exec
DELETE FROM
  main_computer
WHERE
  id = ?;