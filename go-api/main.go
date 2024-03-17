package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const BindingAddr = "localhost:3000"

func main() {
	err := initDatabaseManager()
	if err != nil {
		fmt.Println(err)
		return
	}
	initSensors()

	sensors := getAllSensorIDs()
	for _, sensor := range sensors {
		temp, _ := getTemp(sensor)
		fmt.Printf("Sensor %s has temperature %f \n", sensor, temp)
	}

	initRouter()
}

func initRouter() {
	var router *gin.Engine = gin.Default()
	router.GET("/test/data", getTestData)
	router.GET("/test/:id", getTestId)
	router.POST("/test/data", addTest)
	router.POST("/test/sensor", addTestSensor)
	router.Run(BindingAddr)
}
