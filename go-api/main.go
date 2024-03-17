package main

import (
	"github.com/gin-gonic/gin"
)

const BindingAddr = "localhost:3000"

func main() {
	initDatabase()

	var router *gin.Engine = gin.Default()
	router.GET("/test/data", getTestData)
	router.GET("/test/:id", getTestId)
	router.POST("/test/data", addTest)
	router.POST("/test/sensor", addTestSensor)
	router.Run(BindingAddr)

}
