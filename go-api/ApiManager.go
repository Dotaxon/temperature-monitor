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
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if result, err := updateSensor(sensor); err != nil {
		Log.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	} else if result == 0 {
		err := fmt.Errorf("did not find matching Sensor")
		Log.Println(err)
		context.IndentedJSON(http.StatusNotFound, err)
		return
	}

	Log.Printf("Set Sensor name of %s (ID) to name: %s \n", sensor.Id, sensor.Name)
	context.IndentedJSON(http.StatusCreated, sensor)
}

func getSensor(context *gin.Context) {
	var id string = context.Param("id")

	sensor, err := GetSensorEntry(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	context.IndentedJSON(http.StatusOK, sensor)
}

func getDataEntries(context *gin.Context) {
	Log.Println("Tries to get data entries")
	var body GetDataEntriesBody

	if err := context.BindJSON(&body); err != nil {
		Log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if !body.Interval.IsValid() {
		err := fmt.Errorf("%d is not an valid interval", body.Interval)
		Log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if !ExitsSensor(body.SensorID) {
		err := fmt.Errorf("did not find Sensor with ID: %s ", body.SensorID)
		Log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if body.StartTime > body.EndTime {
		err := fmt.Errorf("startTime is bigger than endTime")
		Log.Println(err)
		context.IndentedJSON(http.StatusBadRequest, err)
	}

	//var dataPoints []SimpleDataPoint
	//var err error
	//switch body.Interval {
	//case Minute:
	//	dataPoints, err = GetEntryCollection(body.StartTime, body.EndTime, body.SensorID, )
	//}

	dataPoints, err := GetEntryCollection(body.StartTime, body.EndTime, body.SensorID, body.Interval)

	if err != nil {
		Log.Println(err)
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, GetDataEntriesResponse{
		SensorID: body.SensorID,
		Data:     dataPoints,
	})
}
