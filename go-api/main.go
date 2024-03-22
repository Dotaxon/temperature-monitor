package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

const BindingAddr = "localhost:3000"
const CertFile = "C:\\Users\\Vincent\\Desktop\\RPI-Heizung_API.crt"
const KeyFile = "C:\\Users\\Vincent\\Desktop\\RPI-Heizung_API.key"

//const CertFile = "/home/vincent/RPI-Heizung.fritz.box.crt"
//const KeyFile = "/home/vincent/RPI-Heizung.fritz.box.key"

var Log *log.Logger

func main() {
	Log = log.Default()
	Log.Println("Logger initialized")

	err := initDatabaseManager()
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Println("Successfully initialized DatabaseManager")

	err = initSensors()
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Println("Successfully initialized Sensors")

	sensors := getRefreshedSensorIDs()

	Log.Printf("Got %d sensors", len(sensors))
	for _, sensor := range sensors {
		temp, _ := getTemp(sensor)
		Log.Printf("Sensor %s has temperature %f \n", sensor, temp)
	}

	//generateEntries(60*1 + 1)

	err = initRouter()
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Println("Successfully initialized Router")
}

func initRouter() error {
	var router *gin.Engine = gin.Default()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200", "http://localhost", "http://RPI-Heizung:4200",
			"http://RPI-Heizung", "http://RPI-Heizung.fritz.box", "https://localhost:4200", "https://localhost", "https://RPI-Heizung:4200",
			"https://RPI-Heizung", "https://RPI-Heizung.fritz.box"},
		AllowedMethods:      []string{"GET", "POST", "PATCH"},
		AllowPrivateNetwork: false,
		AllowCredentials:    true,
		// Enable Debugging for testing, consider disabling in production
		//Debug: true,
	})

	router.Use(c)
	router.PATCH("/sensor/update", updateSensorName)
	router.GET("/sensor/:id", getSensor)
	router.GET("/sensors", getSensors)
	router.POST("/data", getDataEntries)

	router.GET("/test/data", getTestData)
	router.GET("/test/:id", getTestId)
	router.POST("/test/data", addTest)
	router.POST("/test/sensor", addTestSensor)
	//return router.RunTLS(BindingAddr, CertFile, KeyFile)
	return router.Run(BindingAddr)
}
