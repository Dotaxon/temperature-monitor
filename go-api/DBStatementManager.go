package main

import (
	"database/sql"
)

// Parameters: sensorID, name
//
// inserts or ignores
var createSensorStmt *sql.Stmt

// Parameters: time, temperature, sensorID, hour
var createDataEntryStmt *sql.Stmt

// Parameters: hour (min, sec, nsec = 0), average, day
var createHourStmt *sql.Stmt

// Parameters: day (hour, min, sec, nsec = 0), average, startWeekDay
var createDayStmt *sql.Stmt

// Parameters: startDay (hour, min, sec, nsec = 0), average
var createWeekStmt *sql.Stmt

// Parameters: hour (min, sec, nsec = 0)
var existsHourStmt *sql.Stmt

// Parameters: day (hour, min, sec, nsec = 0)
var existsDayStmt *sql.Stmt

// Parameters: mondayOfTheWeek (hour, min, sec, nsec = 0)
var existsWeekStmt *sql.Stmt

// Parameters: toAddTemperature, hour (min, sec, nsec = 0)
//
// The average calc is done by this statement
var updateHourStmt *sql.Stmt

// Parameters: toAddTemperature, day (hour, min, sec, nsec = 0)
//
// The average calc is done by this statement
var updateDayStmt *sql.Stmt

// Parameters: toAddTemperature, startDay (hour, min, sec, nsec = 0)
//
// The average calc is done by this statement
var updateWeekStmt *sql.Stmt

// Parameters: new name, sensorID
var updateSensorNameStmt *sql.Stmt

func initStatements() error {
	//region Create statements
	stmt, err := database.Prepare(`INSERT OR IGNORE INTO sensor (sensorID, name) VALUES (?,?)`)
	if err != nil {
		return err
	}
	createSensorStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO dataEntry (time, temperature, sensorID, hour) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	createDataEntryStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO hourCollection (hour, average, day, numberOfElements) VALUES (?,?,?,1)`)
	if err != nil {
		return err
	}
	createHourStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO dayCollection (day, average, weekStartDay, numberOfElements) VALUES (?,?,?,1)`)
	if err != nil {
		return err
	}
	createDayStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO weekCollection (startDay, average, numberOfElements) VALUES (?,?,1)`)
	if err != nil {
		return err
	}
	createWeekStmt = stmt
	//endregion Create statements

	//region Exists Statements
	stmt, err = database.Prepare(`SELECT 1 FROM hourCollection h WHERE h.hour = ?`)
	if err != nil {
		return err
	}
	existsHourStmt = stmt

	stmt, err = database.Prepare(`SELECT 1 FROM dayCollection d WHERE d.day = ?`)
	if err != nil {
		return err
	}
	existsDayStmt = stmt

	stmt, err = database.Prepare(`SELECT 1 FROM weekCollection w WHERE w.startDay = ?`)
	if err != nil {
		return err
	}
	existsWeekStmt = stmt
	//endregion

	//region UPDATE Statements
	stmt, err = database.Prepare(
		`UPDATE hourCollection SET
                          average = (average * numberOfElements + ?) / (numberOfElements + 1),
                          numberOfElements = numberOfElements + 1
                      WHERE hour = ?`)
	if err != nil {
		return err
	}
	updateHourStmt = stmt

	stmt, err = database.Prepare(
		`UPDATE dayCollection SET
                          average = (average * numberOfElements + ?) / (numberOfElements + 1),
                          numberOfElements = numberOfElements + 1
                      WHERE day = ?`)
	if err != nil {
		return err
	}
	updateDayStmt = stmt

	stmt, err = database.Prepare(
		`UPDATE weekCollection SET
                          average = (average * numberOfElements + ?) / (numberOfElements + 1),
                          numberOfElements = numberOfElements + 1
                      WHERE startDay = ?`)
	if err != nil {
		return err
	}
	updateWeekStmt = stmt

	stmt, err = database.Prepare(`UPDATE sensor SET name = ? WHERE sensorID = ?`)
	if err != nil {
		return err
	}
	updateSensorNameStmt = stmt
	//endregion UPDATE Statements

	return nil
}

func createTables() {

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dataEntry" (
		"entryID"	INTEGER NOT NULL PRIMARY KEY,
		"time"	NUMERIC,
		"temperature"	REAL,
		"sensorID"	TEXT,
		"hour"	NUMERIC,
		FOREIGN KEY("hour") REFERENCES "hourCollection"("hour"),
		FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "hourCollection" (
		"hour"	NUMERIC NOT NULL,
		"average"	REAL,
		"day"	NUMERIC,
		"numberOfElements"	INTEGER,
		PRIMARY KEY("hour"),
		FOREIGN KEY("day") REFERENCES "dayCollection"("day")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dayCollection" (
		"day"	NUMERIC NOT NULL,
		"average"	REAL,
		"weekStartDay"	NUMERIC,
		"numberOfElements"	INTEGER,
		PRIMARY KEY("day"),
		FOREIGN KEY("weekStartDay") REFERENCES "weekCollection"("startDay")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "weekCollection" (
		"startDay"	NUMERIC NOT NULL,
		"average"	REAL,
		"numberOfElements"	INTEGER,
		PRIMARY KEY("startDay")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "sensor" (
		"sensorID"	TEXT NOT NULL,
		"name"	TEXT,
		PRIMARY KEY("sensorID")
	);`)

	Log.Println("Created tables if necessary")
}
