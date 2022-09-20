CREATE TABLE sensors (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  mesure TEXT
);

CREATE TABLE sensors_data (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  value TEXT NOT NULL,
  sensor_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(sensor_id) REFERENCES sensors(id)
);

CREATE TABLE brake_manager (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  state INTEGER NOT NULL,
  hydrolic_pressure_loss INTEGER CHECK(hydrolic_pressure_loss IN (0, 1)) NOT NULL,
  CRITICAL_pod_acceleration_mesure_timeout INTEGER CHECK(
    CRITICAL_pod_acceleration_mesure_timeout IN (0, 1)
  ) NOT NULL,
  CRITICAL_pod_deceleration_instruction_timeout INTEGER CHECK(
    CRITICAL_pod_deceleration_instruction_timeout IN (0, 1)
  ) NOT NULL,
  verin_blocked INTEGER CHECK(verin_blocked IN (0, 1)) NOT NULL,
  emergency_valve_open_without_hydrolic_pressor_diminution INTEGER CHECK(
    emergency_valve_open_without_hydrolic_pressor_diminution IN (0, 1)
  ) NOT NULL,
  CRITICAL_emergency_brakes_without_deceleration INTEGER CHECK(
    CRITICAL_emergency_brakes_without_deceleration IN (0, 1)
  ) NOT NULL,
  mesured_distance_less_than_desired INTEGER CHECK(mesured_distance_less_than_desired IN (0, 1)) NOT NULL,
  mesured_distance_greater_as_desired INTEGER CHECK(mesured_distance_greater_as_desired IN (0, 1)) NOT NULL
);

CREATE TABLE main_computer (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  state INTEGER NOT NULL
);

CREATE TABLE control_data (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  main_computer_id INTEGER NOT NULL,
  brake_manager_id INTEGER NOT NULL,
  value TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY(main_computer_id) REFERENCES main_computer(id),
  FOREIGN KEY (brake_manager_id) REFERENCES brake_manager(id)
);