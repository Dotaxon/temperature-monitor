package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
