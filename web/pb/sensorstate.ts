/* eslint-disable */

export const protobufPackage = "pb";

export interface BrakeManager {
  state: BrakeManager_States;
  hydrolicPressureLoss: boolean;
  criticalPodAccelerationMesureTimeout: boolean;
  criticalPodDecelerationInstructionTimeout: boolean;
  verinBlocked: boolean;
  emergencyValveOpenWithoutHydrolicPressorDiminution: boolean;
  criticalEmergencyBrakesWithoutDeceleration: boolean;
  mesuredDistanceLessThanDesired: boolean;
  mesuredDistanceGreaterAsDesired: boolean;
}

export enum BrakeManager_States {
  INIT = 0,
  IDLE = 1,
  CALIBRATION = 2,
  READY = 3,
  BRAKING = 4,
  EMERGENCY_BRAKING = 5,
  UNRECOGNIZED = -1,
}

export interface MainComputer {
  state: MainComputer_States;
}

export enum MainComputer_States {
  SLEEP = 0,
  STATIC_FAULT = 1,
  DYNAMIC_FAULT = 2,
  SAFE_TO_APPROACH = 3,
  INITIALIZATION = 4,
  SAFE_TO_LAUNCH = 5,
  ACCELERATING = 6,
  EMERGENCY_BRAKING = 7,
  END_BRAKES = 8,
  UNRECOGNIZED = -1,
}

export interface SensorState {
  mainComputer: MainComputer | undefined;
  brakeManager: BrakeManager | undefined;
}
