package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rkusa/gm/math32"
	"log"
	Time "time"
)

var database *sql.DB

func initDatabaseManager() error {
	db, err := sql.Open("sqlite3", DatabasePath)
	if err != nil {
		return err
	}
	database = db
	Log.Println("Successfully opened DB")

	err = createTables()
	if err != nil {
		return err
	}
	Log.Println("Successfully created Tables")

	err = initStatements()
	if err != nil {
		return err
	}
	Log.Println("Successfully initialized Statements")
	return nil
}

// Updates the name of the sensor. The name inside sensor is the new name
//
// Returns the number of affected rows (-1 if an error occurred) and an error
func updateSensor(sensor Sensor) (int64, error) {
	if result, err := updateSensorNameStmt.Exec(sensor.Name, sensor.Id); err != nil {
		return -1, err
	} else if rows, err2 := result.RowsAffected(); err2 != nil {
		return -1, err2
	} else {
		return rows, nil
	}
}

func createSensorIfNotExists(sensor Sensor) error {
	_, err := createSensorStmt.Exec(sensor.Id, sensor.Name)
	if err != nil {
		return err
	}
	return nil
}

func createEntry(time Time.Time, sensorID string, temperature float32) error {
	timeUTC := time.UTC()
	hourTime := getToHourReducedTimeUTC(timeUTC).Unix()
	dayTime := getToDayReducedTimeUTC(timeUTC).Unix()
	weekTime := getToWeekReducedTimeUTC(timeUTC).Unix() //Aka startDay

	if !existsWeekTime(weekTime, sensorID) {
		_, err := createWeekStmt.Exec(weekTime, temperature, sensorID)
		if err != nil {
			return err
		}
	} else {
		_, err := updateWeekStmt.Exec(temperature, weekTime, sensorID)
		if err != nil {
			return err
		}
	}

	if !existsDayTime(dayTime, sensorID) {
		_, err := createDayStmt.Exec(dayTime, temperature, sensorID, weekTime)
		if err != nil {
			return err
		}
	} else {
		_, err := updateDayStmt.Exec(temperature, dayTime, sensorID)
		if err != nil {
			return err
		}
	}

	if !existsHourTime(hourTime, sensorID) {
		_, err := createHourStmt.Exec(hourTime, temperature, sensorID, dayTime)
		if err != nil {
			return err
		}
	} else {
		_, err := updateHourStmt.Exec(temperature, hourTime, sensorID)
		if err != nil {
			return err
		}
	}

	_, err := createDataEntryStmt.Exec(timeUTC.Unix(), temperature, sensorID, hourTime)
	if err != nil {
		return err
	}

	return nil
}

//region Exists

// ExitsSensor Checks if sensor exists
func ExitsSensor(sensorID string) bool {
	rows, err := existsSensorStmt.Query(sensorID)
	if err != nil {
		log.Println(err)
		return false
	}

	if !rows.Next() {
		return false
	}

	err = rows.Close()
	if err != nil {
		log.Println("Unable to close row in existsSensor")
	}
	return true
}

// Expects a to hour reduced Time (like you get from getToHourReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsHourTime(time int64, sensorID string) bool {
	rows, err := existsHourStmt.Query(time, sensorID)
	if err != nil {
		log.Println(err)
		return false
	}

	if !rows.Next() {
		return false
	}

	err = rows.Close()
	if err != nil {
		log.Println("Unable to close row in existsHourTime")
	}
	return true
}

// Expects a to day reduced Time (like you get from getToDayReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsDayTime(time int64, sensorID string) bool {
	rows, err := existsDayStmt.Query(time, sensorID)
	if err != nil {
		return false
	}

	if !rows.Next() {
		return false
	}

	err = rows.Close()
	if err != nil {
		log.Println("Unable to close row in existsDayTime")
	}
	return true
}

// Expects a to week reduced Time (like you get from getToWeekReducedTimeUTC) in Unix format
// and checks if it exists in the hour collection
func existsWeekTime(time int64, sensorID string) bool {
	rows, err := existsWeekStmt.Query(time, sensorID)
	if err != nil {
		return false
	}

	if !rows.Next() {
		return false
	}

	err = rows.Close()
	if err != nil {
		log.Println("Unable to close row in existsWeekTime")
	}
	return true
}

//endregion Exists

//region getToReducedTimeUTC

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

//endregion getToReducedTimeUTC

//region getData

func GetSensorEntry(sensorID string) (Sensor, error) {
	rows, err := getSensorStmt.Query(sensorID)
	defer closeRows(rows)
	if err != nil {
		return Sensor{}, err
	}

	if !rows.Next() {
		return Sensor{}, fmt.Errorf("sensor (%s) not found", sensorID)
	}

	var sensor Sensor
	err = rows.Scan(&sensor.Id, &sensor.Name)
	if err != nil {
		return Sensor{}, err
	}

	return sensor, nil
}

func GetSensorEntries() ([]Sensor, error) {
	rows, err := getSensorsStmt.Query()
	defer closeRows(rows)
	if err != nil {
		return nil, err
	}

	sensors := make([]Sensor, 0, 20)

	for rows.Next() {
		var sensor Sensor
		if err := rows.Scan(&sensor.Id, &sensor.Name); err != nil {
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func GetEntryCollection(startTimeUnix int64, endTimeUnix int64, sensorID string, interval Interval) ([]SimpleDataPoint, error) {
	var startTime Time.Time
	var endTime Time.Time
	var rows *sql.Rows
	var err error
	var difference float64

	defer closeRows(rows)

	switch interval {
	case Minute:
		startTime = Time.Unix(startTimeUnix, 0)
		endTime = Time.Unix(endTimeUnix, 0)
		difference = endTime.Sub(startTime).Minutes()
		rows, err = getDataEntriesFromToOf.Query(startTime.Unix(), endTime.Unix(), sensorID)
	case Hour:
		startTime = getToHourReducedTimeUTC(Time.Unix(startTimeUnix, 0))
		endTime = getToHourReducedTimeUTC(Time.Unix(endTimeUnix, 0))
		difference = endTime.Sub(startTime).Hours()
		rows, err = getHourEntriesFromToOf.Query(startTime.Unix(), endTime.Unix(), sensorID)
	case Day:
		startTime = getToDayReducedTimeUTC(Time.Unix(startTimeUnix, 0))
		endTime = getToDayReducedTimeUTC(Time.Unix(endTimeUnix, 0))
		difference = endTime.Sub(startTime).Hours() / 24
		rows, err = getDayEntriesFromToOf.Query(startTime.Unix(), endTime.Unix(), sensorID)
	case Week:
		startTime = getToWeekReducedTimeUTC(Time.Unix(startTimeUnix, 0))
		endTime = getToWeekReducedTimeUTC(Time.Unix(endTimeUnix, 0))
		difference = endTime.Sub(startTime).Hours() / (24 * 7)
		rows, err = getWeekEntriesFromToOf.Query(startTime.Unix(), endTime.Unix(), sensorID)
	default:
		err = fmt.Errorf("invalid interval %d is not in range", interval)
	}

	if err != nil {
		Log.Println(err)
		return nil, err
	}

	dataPoints := make([]SimpleDataPoint, 0, int64(difference)+1)

	for rows.Next() {
		var time int64 = 0
		var temp float32 = math32.NaN()

		if err := rows.Scan(&time, &temp); err != nil {
			return nil, err
		}
		dataPoints = append(dataPoints, SimpleDataPoint{Time: time, Temp: temp})
	}

	return dataPoints, nil
}

func closeRows(rows *sql.Rows) {
	if rows == nil {
		return
	}

	if err := rows.Close(); err != nil {
		Log.Println(err)
	}
}

//endregion getData
