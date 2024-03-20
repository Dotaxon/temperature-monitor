package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rkusa/gm/math32"
	"github.com/yryz/ds18b20"
	"net/http"
)

var sensors []string

func initSensors() error {
	refreshSensors()

	for _, sensor := range sensors {
		err := createSensorIfNotExists(Sensor{Id: sensor, Name: sensor})
		if err != nil {
			return err
		}
	}

	return nil
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

func updateSensorName(context *gin.Context) {
	var sensor Sensor
	if err := context.BindJSON(&sensor); err != nil {
		Log.Println(err)
		return
	}

	if result, err := updateSensor(sensor); err != nil {
		Log.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, err)
	} else if result == 0 {
		context.IndentedJSON(http.StatusNotFound, nil)
	}

	Log.Printf("Set Sensor name of %s (ID) to name: %s \n", sensor.Id, sensor.Name)
	context.IndentedJSON(http.StatusCreated, sensor)
}
