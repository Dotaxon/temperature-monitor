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
		"hourID"	INTEGER NOT NULL,
		"time"	NUMERIC,
		"average"	REAL,
		"dayID"	INTEGER,
		PRIMARY KEY("hourID"),
		FOREIGN KEY("dayID") REFERENCES "dayCollection"("dayID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "dayCollection" (
		"dayID"	INTEGER NOT NULL,
		"day"	NUMERIC,
		"average"	REAL,
		"weekID"	INTEGER,
		FOREIGN KEY("weekID") REFERENCES "weekCollection"("weekID"),
		PRIMARY KEY("dayID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "weekCollection" (
		"weekID"	INTEGER NOT NULL,
		"startDay"	NUMERIC,
		"average"	REAL,
		PRIMARY KEY("weekID")
	);`)

	_, _ = database.Exec(`CREATE TABLE IF NOT EXISTS "sensor" (
		"sensorID"	TEXT NOT NULL,
		"name"	TEXT,
		PRIMARY KEY("sensorID")
	);`)

	fmt.Println("Created tables if necessary")
}
