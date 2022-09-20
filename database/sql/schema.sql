CREATE TABLE main_computers (
  id integer PRIMARY KEY NOT NULL,
  state integer NOT NULL
);

CREATE TABLE brake_managers (
  id integer PRIMARY KEY,
  state integer NOT NULL,
  hydrolic_pressure_loss integer NOT NULL,
  critical_pod_acceleration_mesure_timeout integer NOT NULL,
  critical_pod_deceleration_instruction_timeout integer NOT NULL,
  verin_blocked integer NOT NULL,
  emergency_valve_open_without_hydrolic_pressor_diminution integer NOT NULL,
  critical_emergency_brakes_without_deceleration integer NOT NULL,
  mesured_distance_less_than_desired integer NOT NULL,
  mesured_distance_greater_as_desired integer NOT NULL
);