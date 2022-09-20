// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: brake_manager.sql

package sqlc

import (
	"context"
)

const createBrakeManager = `-- name: CreateBrakeManager :one
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
    (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, state, hydrolic_pressure_loss, critical_pod_acceleration_mesure_timeout, critical_pod_deceleration_instruction_timeout, verin_blocked, emergency_valve_open_without_hydrolic_pressor_diminution, critical_emergency_brakes_without_deceleration, mesured_distance_less_than_desired, mesured_distance_greater_as_desired
`

type CreateBrakeManagerParams struct {
	State                                              int64 `json:"state"`
	HydrolicPressureLoss                               int64 `json:"hydrolic_pressure_loss"`
	CriticalPodAccelerationMesureTimeout               int64 `json:"critical_pod_acceleration_mesure_timeout"`
	CriticalPodDecelerationInstructionTimeout          int64 `json:"critical_pod_deceleration_instruction_timeout"`
	VerinBlocked                                       int64 `json:"verin_blocked"`
	EmergencyValveOpenWithoutHydrolicPressorDiminution int64 `json:"emergency_valve_open_without_hydrolic_pressor_diminution"`
	CriticalEmergencyBrakesWithoutDeceleration         int64 `json:"critical_emergency_brakes_without_deceleration"`
	MesuredDistanceLessThanDesired                     int64 `json:"mesured_distance_less_than_desired"`
	MesuredDistanceGreaterAsDesired                    int64 `json:"mesured_distance_greater_as_desired"`
}

func (q *Queries) CreateBrakeManager(ctx context.Context, arg CreateBrakeManagerParams) (BrakeManager, error) {
	row := q.db.QueryRowContext(ctx, createBrakeManager,
		arg.State,
		arg.HydrolicPressureLoss,
		arg.CriticalPodAccelerationMesureTimeout,
		arg.CriticalPodDecelerationInstructionTimeout,
		arg.VerinBlocked,
		arg.EmergencyValveOpenWithoutHydrolicPressorDiminution,
		arg.CriticalEmergencyBrakesWithoutDeceleration,
		arg.MesuredDistanceLessThanDesired,
		arg.MesuredDistanceGreaterAsDesired,
	)
	var i BrakeManager
	err := row.Scan(
		&i.ID,
		&i.State,
		&i.HydrolicPressureLoss,
		&i.CriticalPodAccelerationMesureTimeout,
		&i.CriticalPodDecelerationInstructionTimeout,
		&i.VerinBlocked,
		&i.EmergencyValveOpenWithoutHydrolicPressorDiminution,
		&i.CriticalEmergencyBrakesWithoutDeceleration,
		&i.MesuredDistanceLessThanDesired,
		&i.MesuredDistanceGreaterAsDesired,
	)
	return i, err
}

const deleteBrakeManager = `-- name: DeleteBrakeManager :exec
DELETE FROM
    brake_manager
WHERE
    id = ?
`

func (q *Queries) DeleteBrakeManager(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBrakeManager, id)
	return err
}

const getBrakeManager = `-- name: GetBrakeManager :one
SELECT
    id, state, hydrolic_pressure_loss, critical_pod_acceleration_mesure_timeout, critical_pod_deceleration_instruction_timeout, verin_blocked, emergency_valve_open_without_hydrolic_pressor_diminution, critical_emergency_brakes_without_deceleration, mesured_distance_less_than_desired, mesured_distance_greater_as_desired
FROM
    brake_manager
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetBrakeManager(ctx context.Context, id int64) (BrakeManager, error) {
	row := q.db.QueryRowContext(ctx, getBrakeManager, id)
	var i BrakeManager
	err := row.Scan(
		&i.ID,
		&i.State,
		&i.HydrolicPressureLoss,
		&i.CriticalPodAccelerationMesureTimeout,
		&i.CriticalPodDecelerationInstructionTimeout,
		&i.VerinBlocked,
		&i.EmergencyValveOpenWithoutHydrolicPressorDiminution,
		&i.CriticalEmergencyBrakesWithoutDeceleration,
		&i.MesuredDistanceLessThanDesired,
		&i.MesuredDistanceGreaterAsDesired,
	)
	return i, err
}

const getBrakeManagers = `-- name: GetBrakeManagers :many
SELECT
    id, state, hydrolic_pressure_loss, critical_pod_acceleration_mesure_timeout, critical_pod_deceleration_instruction_timeout, verin_blocked, emergency_valve_open_without_hydrolic_pressor_diminution, critical_emergency_brakes_without_deceleration, mesured_distance_less_than_desired, mesured_distance_greater_as_desired
FROM
    brake_manager
ORDER BY
    id
`

func (q *Queries) GetBrakeManagers(ctx context.Context) ([]BrakeManager, error) {
	rows, err := q.db.QueryContext(ctx, getBrakeManagers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BrakeManager
	for rows.Next() {
		var i BrakeManager
		if err := rows.Scan(
			&i.ID,
			&i.State,
			&i.HydrolicPressureLoss,
			&i.CriticalPodAccelerationMesureTimeout,
			&i.CriticalPodDecelerationInstructionTimeout,
			&i.VerinBlocked,
			&i.EmergencyValveOpenWithoutHydrolicPressorDiminution,
			&i.CriticalEmergencyBrakesWithoutDeceleration,
			&i.MesuredDistanceLessThanDesired,
			&i.MesuredDistanceGreaterAsDesired,
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
