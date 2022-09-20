-- name: GetMainComputer :one
SELECT
  *
FROM
  main_computers
WHERE
  id = ?
LIMIT
  1;

-- name: GetMainComputerCount :one
SELECT
  COUNT(*)
FROM
  main_computers;

-- name: GetMainComputers :many
SELECT
  *
FROM
  main_computers
ORDER BY
  id;

-- name: CreateMainComputer :one
INSERT INTO
  main_computers (state)
VALUES
  (?) RETURNING *;

-- name: DeleteMainComputer :exec
DELETE FROM
  main_computers
WHERE
  id = ?;