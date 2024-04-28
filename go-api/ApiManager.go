package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func updateSensorName(context *gin.Context) {
	Log.Println("Tries to update Sensor")
	var sensor Sensor
	if err := context.BindJSON(&sensor); err != nil {
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if result, err := updateSensor(sensor); err != nil {
		Log.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	} else if result == 0 {
		err := fmt.Errorf("did not find matching Sensor")
		Log.Println(err)
		context.JSON(http.StatusNotFound, err)
		return
	}

	Log.Printf("Set Sensor name of %s (ID) to name: %s \n", sensor.Id, sensor.Name)
	context.JSON(http.StatusCreated, sensor)
}

func getSensor(context *gin.Context) {
	var id = context.Param("id")

	sensor, err := GetSensorEntry(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, sensor)
}

func getSensors(context *gin.Context) {
	sensors, err := GetSensorEntries()
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, sensors)
}

func getDataEntries(context *gin.Context) {
	Log.Println("Tries to get data entries")
	var body GetDataEntriesBody

	if err := context.BindJSON(&body); err != nil {
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if !body.Interval.IsValid() {
		err := fmt.Errorf("%d is not an valid interval", body.Interval)
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if !ExitsSensor(body.SensorID) {
		err := fmt.Errorf("did not find Sensor with ID: %s ", body.SensorID)
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if body.StartTime > body.EndTime {
		err := fmt.Errorf("startTime is bigger than endTime")
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
	}

	dataPoints, err := GetEntryCollection(body.StartTime, body.EndTime, body.SensorID, body.Interval)

	if err != nil {
		Log.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, GetDataEntriesResponse{
		SensorID: body.SensorID,
		Data:     dataPoints,
	})
}

func getSensorsWithTemp(context *gin.Context) {
	var sensors []Sensor

	if err := context.BindJSON(&sensors); err != nil {
		Log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	var sensorsT []SensorWithTemp
	sensorsT = getTempsFrom(sensors)
	if len(sensorsT) != 0 {
		Log.Println("Got no temperatures at all")
		context.JSON(http.StatusInternalServerError, "Got no temperatures at all")
		return
	}

	context.JSON(http.StatusOK, sensorsT)
}
