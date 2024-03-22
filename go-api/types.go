package main

import (
	"time"
)

//region Interval type

type Interval int

const (
	Minute Interval = 1
	Hour   Interval = 2
	Day    Interval = 3
	Week   Interval = 4
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

func (i Interval) IsValid() bool {
	return i == Minute || i == Hour || i == Day || i == Week
}

//endregion Interval type

type Sensor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SensorWithTemp struct {
	Sensor Sensor  `json:"sensor"`
	Temp   float32 `json:"temp"`
}

type DataPoint struct {
	Time time.Time `json:"time"`
	Temp float32   `json:"temp"`
}

type SimpleDataPoint struct {
	Time int64   `json:"time"`
	Temp float32 `json:"temp"`
}

type DataCollection struct {
	Time    time.Time   `json:"time"`
	Average float32     `json:"average"`
	Data    []DataPoint `json:"data"`
}

// GetDataEntriesBody Body of GET request to get data from StartTime to EndTime in specified Interval from SensorID
//
// StartTime and EndTime are in UNIX UTC format and are not reduced to their Interval like getToHourReducedTimeUTC would do
type GetDataEntriesBody struct {
	StartTime int64    `json:"startTime"`
	EndTime   int64    `json:"endTime"`
	SensorID  string   `json:"sensorID"`
	Interval  Interval `json:"interval"`
}

type GetDataEntriesResponse struct {
	SensorID string            `json:"sensorID"`
	Data     []SimpleDataPoint `json:"data"`
}
