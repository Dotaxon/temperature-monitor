package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
	"os"
)

const BindingAddr = "localhost:3000"
const CertFile = "C:\\Users\\Vincent\\Desktop\\RPI-Heizung.fritz.box.crt"
const KeyFile = "C:\\Users\\Vincent\\Desktop\\RPI-Heizung.fritz.box.key"
const DatabasePath = "./data.db"

//const BindingAddr = "RPI-Heizung.fritz.box:3000"
//const CertFile = "/etc/ssl/certs/RPI-Heizung.fritz.box.chained.crt"
//const KeyFile = "/etc/ssl/private/RPI-Heizung.fritz.box.key"
//const DatabasePath = "/etc/GO-API/data.db"

const UseWebSensors = true

var Log *log.Logger

func main() {
	//Log = log.Default()
	Log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	Log.Println("Logger initialized")

	err := initDatabaseManager()
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Println("Successfully initialized DatabaseManager")

	initWebSensors()

	err = initSensors()
	if err != nil {
		Log.Fatal(err)
		return
	}
	Log.Println("Successfully initialized Sensors")

	sensors := getRefreshedSensorIDs()

	Log.Printf("Got %d sensors", len(sensors))
	for _, sensor := range sensors {
		temp, err := getTemp(sensor)
		if err != nil {
			Log.Printf("Error while reading temp %s\n", err.Error())
		} else {
			Log.Printf("Sensor %s has temperature %f \n", sensor, temp)
		}
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
			"https://RPI-Heizung", "https://RPI-Heizung.fritz.box", "https://RPI-Heizung.olbring.org", "https://192.168.178.10",
			"http://RPI-Heizung.olbring.org", "http://192.168.178.10"},
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
	router.POST("/sensors/temps", getSensorsWithTemp)

	router.GET("/test/data", getTestData)
	router.GET("/test/:id", getTestId)
	router.POST("/test/data", addTest)
	router.POST("/test/sensor", addTestSensor)
	//return router.RunTLS(BindingAddr, CertFile, KeyFile)
	return router.Run(BindingAddr)
}
