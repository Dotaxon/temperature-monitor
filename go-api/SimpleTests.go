package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
		Log.Println("Ah error")
		return
	}

	Log.Printf("time: %s temp: %f", dataPoint.Time.String(), dataPoint.Temp)

	context.IndentedJSON(http.StatusCreated, dataPoint)
}

func addTestSensor(context *gin.Context) {
	var sensor Sensor

	if err := context.BindJSON(&sensor); err != nil {
		//error handling
		Log.Println(err)
		Log.Println("Ah error")
		return
	}

	Log.Printf("Id: %s Name: %s \n", sensor.Id, sensor.Name)

	context.IndentedJSON(http.StatusCreated, sensor)
}

func generateEntries(amount int) {
	startTime := Time.Now()
	time := Time.Date(2001, 1, 1, 0, 0, 0, 0, Time.UTC)

	sensors := getAllSensorIDs()

	if len(sensors) == 0 {
		sensors = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
		Log.Printf("Used artifical sensors")
	}

	for range amount {
		for _, sensor := range sensors {
			err := createEntry(time, sensor, float32(rand.Int32N(31)))
			if err != nil {
				log.Println(err)
			}
		}
		time = time.Add(Time.Minute)
	}
	neededTime := Time.Since(startTime).Seconds()
	entries := amount * len(sensors)
	log.Printf("Created %d entries in %fs with an average of %fs/entrie", entries, neededTime, neededTime/float64(entries))
}
