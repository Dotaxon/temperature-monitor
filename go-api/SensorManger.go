package main

import (
	"fmt"
	"github.com/rkusa/gm/math32"
	"github.com/yryz/ds18b20"
	"strings"
	"sync"
	Time "time"
)

var sensors []string
var sensorsMutex = sync.RWMutex{}
var OUTSIDE_TEMP_SENSOR = "OUTSIDE_TEMP_SENSOR"

func initSensors() error {
	refreshSensors()

	sensorsMutex.RLock()
	defer sensorsMutex.RUnlock()
	for _, sensor := range sensors {
		err := createSensorIfNotExists(Sensor{Id: sensor, Name: sensor})
		if err != nil {
			return err
		}
	}

	go sensorTrigger()

	return nil
}

func sensorTrigger() {
	for range Time.Tick(Time.Minute) {
		go createCompleteMeasurement()
	}
}

func createCompleteMeasurement() {
	refreshSensors()

	Log.Println("Tick")

	sensorsMutex.RLock()
	sensorList := make([]string, len(sensors))
	copy(sensorList, sensors)
	sensorsMutex.RUnlock()

	time := Time.Now()

	for _, sensor := range sensorList {
		temp, err := getTemp(sensor)
		if err != nil {
			Log.Println(err)
			continue
		}

		err = createEntry(time, sensor, temp)
		if err != nil {
			Log.Println(err)
			continue
		}
	}
	Log.Println("Tack")
}

func refreshSensors() {
	newSensorCandidates, _ := ds18b20.Sensors()
	newSensors := make([]string, 0, len(newSensorCandidates)+1)

	for _, sensorCandidate := range newSensorCandidates {
		if strings.HasPrefix(sensorCandidate, "28") {
			newSensors = append(newSensors, sensorCandidate)
		}
	}

	if UseOutsideSensor {
		newSensors = append(newSensors, OUTSIDE_TEMP_SENSOR)
	}

	sensorsMutex.Lock()
	defer sensorsMutex.Unlock()
	sensors = newSensors
}

func getRefreshedSensorIDs() []string {
	refreshSensors()

	sensorsMutex.RLock()
	defer sensorsMutex.RUnlock()

	sensorList := make([]string, len(sensors))
	copy(sensorList, sensors)

	return sensorList
}

func getSensorTemp(sensor Sensor) (float32, error) {
	return getTemp(sensor.Id)
}

func getTemp(sensorID string) (float32, error) {

	//Uses special behaviour for outside temp sensor
	if sensorID == OUTSIDE_TEMP_SENSOR {
		temp, err := GetTempViaWeb()
		if err != nil {
			return math32.NaN(), fmt.Errorf("could not read %s error: %s", sensorID, err.Error())
		}
		return temp, nil
	}

	temp, err := ds18b20.Temperature(sensorID)
	if err != nil {
		return math32.NaN(), fmt.Errorf("could not read %s", sensorID)
	}
	return float32(temp), nil
}

func getTempsFrom(sensors []Sensor) []SensorWithTemp {
	list := make([]SensorWithTemp, 0, len(sensors))

	for _, sensor := range sensors {
		temp, err := getSensorTemp(sensor)
		if err != nil {
			Log.Println(err)
			temp = -999
		}

		list = append(list, SensorWithTemp{
			Sensor: sensor,
			Temp:   temp,
		})
	}
	return list
}
