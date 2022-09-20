-- name: GetBrakeManager :one
SELECT
    *
FROM
    brake_manager
WHERE
    id = ?
LIMIT
    1;

-- name: GetBrakeManagers :many
SELECT
    *
FROM
    brake_manager
ORDER BY
    id;

-- name: CreateBrakeManager :one
INSERT INTO
    brake_manager (
        state,
        hydrolic_pressure_loss,
        critical_pod_acceleration_mesure_timeout,
        critical_pod_deceleration_instruction_timeout,
        verin_blocked,
        emergency_valve_open_without_hydrolic_pressor_diminution,
        critical_emergency_brakes_without_deceleration,
        mesured_distance_less_than_desired,
        mesured_distance_greater_as_desired
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: DeleteBrakeManager :exec
DELETE FROM
    brake_manager
WHERE
    id = ?;