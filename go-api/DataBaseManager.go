package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	Time "time"
)

var database *sql.DB

// Parameters: sensorID, name
var createSensorStmt *sql.Stmt

// Parameters: time, temperature, sensorID, hour
var createDataEntryStmt *sql.Stmt

// Parameters: hour, average, day
var createHourStmt *sql.Stmt

// Parameters: day, average, startWeekDay
var createDayStmt *sql.Stmt

// Parameters: startDay, average
var createWeekStmt *sql.Stmt

// Parameters: hour (min, sec, nsec = 0)
var existsHourStmt *sql.Stmt

// Parameters: day (hour, min, sec, nsec = 0)
var existsDayStmt *sql.Stmt

// Parameters: mondayOfTheWeek (hour, min, sec, nsec = 0)
var existsWeekStmt *sql.Stmt

func initDatabaseManager() error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	database = db

	createTables()
	err = initStatements()
	if err != nil {
		return err
	}

	return nil
}

func createEntry(sensorID string, temperature float32) error {
	timeNowUTC := Time.Now().UTC()
	hourTime := getToHourReducedTimeUTC(timeNowUTC).Unix()
	dayTime := getToDayReducedTimeUTC(timeNowUTC).Unix()
	weekTime := getToWeekReducedTimeUTC(timeNowUTC).Unix()

	if !existsWeekTime(weekTime) {
		_, err := createWeekStmt.Exec(weekTime, temperature)
		if err != nil {
			return err
		}
	} else {
		//TODO add updates for week day hour time
	}

	if !existsDayTime(dayTime) {
		_, err := createDayStmt.Exec(dayTime, temperature, weekTime)
		if err != nil {
			return err
		}
	}

	if !existsHourTime(hourTime) {
		_, err := createHourStmt.Exec(hourTime, temperature, dayTime)
		if err != nil {
			return err
		}
	}

	return nil
}

// Expects a to hour reduced Time (like you get from getToHourReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsHourTime(time int64) bool {
	row, err := existsHourStmt.Query(time)
	if err != nil {
		return false
	}
	return row.Next()
}

// Expects a to day reduced Time (like you get from getToDayReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsDayTime(time int64) bool {
	row, err := existsDayStmt.Query(time)
	if err != nil {
		return false
	}
	return row.Next()
}

// Expects a to week reduced Time (like you get from getToWeekReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsWeekTime(time int64) bool {
	row, err := existsWeekStmt.Query(time)
	if err != nil {
		return false
	}
	return row.Next()
}

// Reduces a Time to their hour means min, sec and nsec are 0
// also sets time to UTC
func getToHourReducedTimeUTC(time Time.Time) Time.Time {
	timeUTC := time.UTC()
	hour := timeUTC.Hour()
	year, month, day := timeUTC.Date()
	return Time.Date(year, month, day, hour, 0, 0, 0, Time.UTC)
}

// Reduces a Time to their day means hour, min, sec and nsec are 0
// also sets time to UTC
func getToDayReducedTimeUTC(time Time.Time) Time.Time {
	year, month, day := time.UTC().Date()
	return Time.Date(year, month, day, 0, 0, 0, 0, Time.UTC)
}

// Reduces a Time to their week (start on Monday) means hour, min, sec and nsec are 0
// and year, month and day are set to the previous monday
// also sets time to UTC
func getToWeekReducedTimeUTC(time Time.Time) Time.Time {
	time = time.UTC()
	for time.Weekday() != Time.Monday {
		time = time.Add(Time.Minute * 60 * 24 * -1) //subtracts a day
	}
	year, month, day := time.Date()
	return Time.Date(year, month, day, 0, 0, 0, 0, Time.UTC)
}

func initStatements() error {
	//region Create statements
	stmt, err := database.Prepare(`INSERT INTO sensor (sensorID, name) VALUES (?,?)`)
	if err != nil {
		return err
	}
	createSensorStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO dataEntry (time, temperature, sensorID, hour) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	createDataEntryStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO hourCollection (hour, average, day) VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	createHourStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO dayCollection (day, average, weekStartDay) VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	createDayStmt = stmt

	stmt, err = database.Prepare(`INSERT INTO weekCollection (startDay, average) VALUES (?,?)`)
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

	return nil
}

func createTables() {
	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dataEntry" (
		"entryID"	INTEGER NOT NULL,
		"time"	NUMERIC,
		"temperature"	REAL,
		"sensorID"	INTEGER,
		"hour"	INTEGER,
		FOREIGN KEY("hour") REFERENCES "hourCollection"("hour"),
		FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID"),
		PRIMARY KEY("entryID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "hourCollection" (
		"hour"	NUMERIC NOT NULL,
		"average"	REAL,
		"day"	INTEGER,
		PRIMARY KEY("hour"),
		FOREIGN KEY("day") REFERENCES "dayCollection"("day")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dayCollection" (
		"day"	NUMERIC NOT NULL,
		"average"	REAL,
		"weekStartDay"	INTEGER,
		PRIMARY KEY("day"),
		FOREIGN KEY("weekStartDay") REFERENCES "weekCollection"("startDay")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "weekCollection" (
		"startDay"	NUMERIC NOT NULL,
		"average"	REAL,
		PRIMARY KEY("startDay")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "sensor" (
		"sensorID"	TEXT NOT NULL,
		"name"	TEXT,
		PRIMARY KEY("sensorID")
	);`)

	fmt.Println("Created tables if necessary")
}
