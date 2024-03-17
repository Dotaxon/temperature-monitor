package main

import (
	"fmt"
	"github.com/rkusa/gm/math32"
	"github.com/yryz/ds18b20"
)

var sensors []string

func initSensors() {
	refreshSensors()
}

func refreshSensors() {
	sensors, _ = ds18b20.Sensors()
}

func getAllSensorIDs() []string {
	refreshSensors()
	return sensors
}

func getSensorTemp(sensor Sensor) (float32, error) {
	return getTemp(sensor.Name)
}

func getTemp(sensor string) (float32, error) {
	temp, err := ds18b20.Temperature(sensor)
	if err != nil {
		return math32.NaN(), fmt.Errorf("could not read %s", sensor)
	}
	return float32(temp), nil
}

func getTempsFrom(sensors []Sensor) ([]SensorWithTemp, error) {
	list := make([]SensorWithTemp, len(sensors))

	for _, sensor := range sensors {
		temp, err := getSensorTemp(sensor)
		if err != nil {
			return nil, err
		}

		list = append(list, SensorWithTemp{
			Sensor: sensor,
			Temp:   temp,
		})
	}
	return list, nil
}
