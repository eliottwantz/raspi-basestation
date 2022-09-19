package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type PolyloopBD struct {
	Db *sql.DB
}

func OpenDB() (*PolyloopBD, error) {
	db, err := sql.Open("sqlite", "polyloop.db")
	if err != nil {
		return nil, err
	}
	return &PolyloopBD{
		Db: db,
	}, nil
}

func (pdb *PolyloopBD) CreateTables() error {
	execString := `
		CREATE TABLE IF NOT EXISTS sensor (
		sensorid INTEGER NOT NULL PRIMARY KEY,
		sensorname TEXT NOT NULL,
		sensormesure TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS sensor_data (
		dataid INTEGER NOT NULL PRIMARY KEY,
		datatime TEXT NOT NULL,
		sensordata TEXT NOT NULL,
		sensorid INTEGER NOT NULL,
		FOREIGN KEY(sensorid) REFERENCES sensor(sensorid)
	);
	CREATE TABLE IF NOT EXISTS control_data (
		controlid INTEGER NOT NULL PRIMARY KEY,
		controltime TEXT NOT NULL,
		maincomputer INTEGER NOT NULL,
		brakemanager INTEGER NOT NULL
		FOREIGN KEY(sensorid) REFERENCES sensor(sensorid)
	);
	`

	pdb, err := OpenDB()
	if err != nil {
		return err
	}

	_, err = pdb.Db.Exec(execString)
	if err != nil {
		return err
	}
	return nil
}
