package main

import (
	"database/sql"
)

//region createStmts

// Parameters: sensorID, name
//
// inserts or ignores
var createSensorStmt *sql.Stmt

// Parameters: time, temperature, sensorID, hour
var createDataEntryStmt *sql.Stmt

// Parameters: hour (min, sec, nsec = 0), average, sensorID, day
var createHourStmt *sql.Stmt

// Parameters: day (hour, min, sec, nsec = 0), average, sensorID, startWeekDay
var createDayStmt *sql.Stmt

// Parameters: startDay (hour, min, sec, nsec = 0), average, sensorID
var createWeekStmt *sql.Stmt

//endregion createStmts

//region existsStmts

// Parameters: hour (min, sec, nsec = 0), sensorID
var existsHourStmt *sql.Stmt

// Parameters: day (hour, min, sec, nsec = 0), sensorID
var existsDayStmt *sql.Stmt

// Parameters: mondayOfTheWeek (hour, min, sec, nsec = 0), sensorID
var existsWeekStmt *sql.Stmt

//endregion existsStmts

//region updateStmts

// Parameters: toAddTemperature, hour (min, sec, nsec = 0), sensorID
//
// The average calc is done by this statement
var updateHourStmt *sql.Stmt

// Parameters: toAddTemperature, day (hour, min, sec, nsec = 0), sensorID
//
// The average calc is done by this statement
var updateDayStmt *sql.Stmt

// Parameters: toAddTemperature, startDay (hour, min, sec, nsec = 0), sensorID
//
// The average calc is done by this statement
var updateWeekStmt *sql.Stmt

// Parameters: new name, sensorID
var updateSensorNameStmt *sql.Stmt

//endregion updateStmts

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

	stmt, err = database.Prepare(`INSERT INTO hourCollection (hour, average, sensorID, day, numberOfElements) VALUES (?,?,?,?,1)`)
	if err != nil {
		return err
	}
	createHourStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO dayCollection (day, average, sensorID, weekStartDay, numberOfElements) VALUES (?,?,?,?,1)`)
	if err != nil {
		return err
	}
	createDayStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO weekCollection (startDay, average, sensorID, numberOfElements) VALUES (?,?,?,1)`)
	if err != nil {
		return err
	}
	createWeekStmt = stmt
	//endregion Create statements

	//region Exists Statements
	stmt, err = database.Prepare(`SELECT 1 FROM hourCollection h WHERE h.hour = ? AND h.sensorID = ?`)
	if err != nil {
		return err
	}
	existsHourStmt = stmt

	stmt, err = database.Prepare(`SELECT 1 FROM dayCollection d WHERE d.day = ? AND d.sensorID = ?`)
	if err != nil {
		return err
	}
	existsDayStmt = stmt

	stmt, err = database.Prepare(`SELECT 1 FROM weekCollection w WHERE w.startDay = ? AND w.sensorID = ?`)
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
                      WHERE hour = ? AND sensorID = ?`)
	if err != nil {
		return err
	}
	updateHourStmt = stmt

	stmt, err = database.Prepare(
		`UPDATE dayCollection SET
                          average = (average * numberOfElements + ?) / (numberOfElements + 1),
                          numberOfElements = numberOfElements + 1
                      WHERE day = ? AND sensorID = ?`)
	if err != nil {
		return err
	}
	updateDayStmt = stmt

	stmt, err = database.Prepare(
		`UPDATE weekCollection SET
                          average = (average * numberOfElements + ?) / (numberOfElements + 1),
                          numberOfElements = numberOfElements + 1
                      WHERE startDay = ? AND sensorID = ?`)
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

var createAllTablesStmt string = `
CREATE TABLE IF NOT EXISTS "sensor" (
	"sensorID"	TEXT NOT NULL,
	"name"	TEXT,
	PRIMARY KEY("sensorID")
);

CREATE TABLE IF NOT EXISTS "dataEntry" (
	"entryID"	INTEGER NOT NULL,
	"time"	NUMERIC,
	"temperature"	REAL,
	"sensorID"	TEXT,
	"hour"	NUMERIC,
	PRIMARY KEY("entryID"),
	FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID"),
	FOREIGN KEY("hour") REFERENCES "hourCollection"("hour")
);

CREATE TABLE IF NOT EXISTS "hourCollection" (
	"hour"	NUMERIC NOT NULL,
	"average"	REAL,
	"day"	NUMERIC,
	"sensorID"	TEXT NOT NULL,
	"numberOfElements"	INTEGER,
	PRIMARY KEY("hour","sensorID"),
	FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID"),
	FOREIGN KEY("day") REFERENCES "dayCollection"("day")
);

CREATE TABLE IF NOT EXISTS "dayCollection" (
	"day"	NUMERIC NOT NULL,
	"average"	REAL,
	"weekStartDay"	NUMERIC,
	"sensorID"	TEXT NOT NULL,
	"numberOfElements"	INTEGER,
	PRIMARY KEY("day","sensorID"),
	FOREIGN KEY("weekStartDay") REFERENCES "weekCollection"("startDay"),
	FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID")
);

CREATE TABLE IF NOT EXISTS "weekCollection" (
	"startDay"	NUMERIC NOT NULL,
	"average"	REAL,
	"sensorID"	INTEGER NOT NULL,
	"numberOfElements"	INTEGER,
	FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID"),
	PRIMARY KEY("startDay","sensorID")
);
`

func createTables() error {
	_, err := database.Exec(createAllTablesStmt)
	return err
}
