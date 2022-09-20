package database

import (
	"app/pb"
)

// type BrakeManager struct {
// 	ID                                                 int64 `json:"id"`
// 	State                                              int64 `json:"state"`
// 	HydrolicPressureLoss                               int64 `json:"hydrolic_pressure_loss"`
// 	CriticalPodAccelerationMesureTimeout               int64 `json:"critical_pod_acceleration_mesure_timeout"`
// 	CriticalPodDecelerationInstructionTimeout          int64 `json:"critical_pod_deceleration_instruction_timeout"`
// 	VerinBlocked                                       int64 `json:"verin_blocked"`
// 	EmergencyValveOpenWithoutHydrolicPressorDiminution int64 `json:"emergency_valve_open_without_hydrolic_pressor_diminution"`
// 	CriticalEmergencyBrakesWithoutDeceleration         int64 `json:"critical_emergency_brakes_without_deceleration"`
// 	MesuredDistanceLessThanDesired                     int64 `json:"mesured_distance_less_than_desired"`
// 	MesuredDistanceGreaterAsDesired                    int64 `json:"mesured_distance_greater_as_desired"`
// }

// type ControlData struct {
// 	ID             int64     `json:"id"`
// 	MainComputerID int64     `json:"main_computer_id"`
// 	BrakeManagerID int64     `json:"brake_manager_id"`
// 	Value          string    `json:"value"`
// 	CreatedAt      time.Time `json:"created_at"`
// }

type MainComputer struct {
	ID uint
	pb.MainComputer
}

type BrakeManager struct {
	ID uint
	pb.BrakeManager
}

// type Sensor struct {
// 	ID     int64  `json:"id"`
// 	Name   string `json:"name"`
// 	Mesure string `json:"mesure"`
// }

// type SensorsData struct {
// 	ID        int64     `json:"id"`
// 	Value     string    `json:"value"`
// 	SensorID  int64     `json:"sensor_id"`
// 	CreatedAt time.Time `json:"created_at"`
// }
