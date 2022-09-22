CREATE TABLE IF NOT EXISTS main_computers (
    id integer NOT NULL,
    created_at datetime NOT NULL,
    state integer NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS brake_managers (
    id integer NOT NULL,
    created_at datetime NOT NULL,
    state integer NOT NULL,
    hydrolic_pressure_loss integer NOT NULL,
    critical_pod_acceleration_mesure_timeout integer NOT NULL,
    critical_pod_deceleration_instruction_timeout integer NOT NULL,
    verin_blocked integer NOT NULL,
    emergency_valve_open_without_hydrolic_pressor_diminution integer NOT NULL,
    critical_emergency_brakes_without_deceleration integer NOT NULL,
    mesured_distance_less_than_desired integer NOT NULL,
    mesured_distance_greater_as_desired integer NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS sensors (
    id integer NOT NULL,
    name text NOT NULL,
    mesure text NOT NULL,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS sensor_datas (
    id integer NOT NULL,
    created_at datetime NOT NULL,
    sensor_id integer NOT NULL,
    value real NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (sensor_id) REFERENCES sensors(id)
);