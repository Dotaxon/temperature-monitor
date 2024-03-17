package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
	Time "time"
)

func getTestData(context *gin.Context) {

	dataPoints := DatePoint{
		Time: Time.Now().UTC(),
		Temp: rand.Float32() * 100,
	}

	context.IndentedJSON(200, dataPoints)
}

func getTestId(context *gin.Context) {
	id := context.Param("id")

	context.IndentedJSON(http.StatusOK, gin.H{"message": id})
}

func addTest(context *gin.Context) {
	var dataPoint DatePoint

	if err := context.BindJSON(&dataPoint); err != nil {
		//error handling
		fmt.Println("Ah error")
		return
	}

	fmt.Printf("time: %s temp: %f", dataPoint.Time.String(), dataPoint.Temp)

	context.IndentedJSON(http.StatusCreated, dataPoint)
}

func addTestSensor(context *gin.Context) {
	var sensor Sensor

	if err := context.BindJSON(&sensor); err != nil {
		//error handling
		fmt.Println(err)
		fmt.Println("Ah error")
		return
	}

	fmt.Printf("Id: %s Name: %s \n", sensor.Id, sensor.Name)

	context.IndentedJSON(http.StatusCreated, sensor)
}
