package main

import "time"

//region Interval type

type Interval int

const (
	Minute Interval = iota
	Hour
	Day
	Week
)

func (i Interval) String() string {
	switch i {
	case Minute:
		return "Minute"
	case Hour:
		return "Hour"
	case Day:
		return "Day"
	case Week:
		return "Week"
	default:
		return "Unknown"
	}
}

//endregion Interval type

type Sensor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SensorWithTemp struct {
	Sensor
	Temp float32 `json:"temp"`
}

type DatePoint struct {
	Time time.Time `json:"time"`
	Temp float32   `json:"temp"`
}

type DataCollection struct {
	Time    time.Time   `json:"time"`
	Average float32     `json:"average"`
	Data    []DatePoint `json:"data"`
}

// GetDataEntriesBody Body of GET request to get data from StartTime to EndTime in specified Interval from SensorID
//
// StartTime and EndTime are in UNIX format and are not reduced to their Interval like getToHourReducedTimeUTC would do
type GetDataEntriesBody struct {
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	SensorID  string `json:"sensorID"`
	Interval  string `json:"interval"`
}
