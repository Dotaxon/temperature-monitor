package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func initDatabase() {
	var db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	database = db
	createTables()
}

func createEntry() {

}

func createTables() {
	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dataEntry" (
		"entryID"	INTEGER NOT NULL,
		"time"	NUMERIC,
		"temperature"	REAL,
		"sensorID"	INTEGER,
		"hourID"	INTEGER,
		FOREIGN KEY("hourID") REFERENCES "hourCollection"("hourID"),
		FOREIGN KEY("sensorID") REFERENCES "sensor"("sensorID"),
		PRIMARY KEY("entryID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "hourCollection" (
		"hour"	NUMERIC NOT NULL,
		"average"	REAL,
		"dayID"	INTEGER,
		PRIMARY KEY("hour"),
		FOREIGN KEY("dayID") REFERENCES "dayCollection"("dayID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dayCollection" (
		"day"	NUMERIC NOT NULL,
		"average"	REAL,
		"weekID"	INTEGER,
		PRIMARY KEY("day"),
		FOREIGN KEY("weekID") REFERENCES "weekCollection"("weekID")
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
