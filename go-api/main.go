package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

const BindingAddr = "localhost:3000"

var Log *log.Logger

func main() {
	Log = log.Default()
	Log.Println("Logger initialized")

	err := initDatabaseManager()
	if err != nil {
		Log.Fatal(err)
		return
	}
	err = initSensors()
	if err != nil {
		Log.Fatal(err)
		return
	}

	sensors := getAllSensorIDs()
	for _, sensor := range sensors {
		temp, _ := getTemp(sensor)
		Log.Printf("Sensor %s has temperature %f \n", sensor, temp)
	}

	generateEntries(60*24*7 + 1)

	err = initRouter()
	if err != nil {
		Log.Fatal(err)
		return
	}
}

func initRouter() error {
	var router *gin.Engine = gin.Default()
	router.GET("/test/data", getTestData)
	router.GET("/test/:id", getTestId)
	router.POST("/test/data", addTest)
	router.POST("/test/sensor", addTestSensor)
	return router.Run(BindingAddr)
}
