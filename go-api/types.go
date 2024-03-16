package main

import "time"

type Sensor struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
