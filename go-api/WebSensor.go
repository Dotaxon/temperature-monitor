package main

import (
	"errors"
	"fmt"
	"github.com/rkusa/gm/math32"
)

// WebSensor represents a sensor that retrieves temperature data via http(s).
type WebSensor struct {
	SensorID        string                  // SensorID represents the unique identifier of the sensor.
	GetTempFunction func() (float32, error) // GetTempFunction is a function reference to retrieve temperature data.
}

// NewWebSensor creates a new WebSensor instance.
func NewWebSensor(sensorID string, getTempFunction func() (float32, error)) *WebSensor {
	return &WebSensor{
		SensorID:        sensorID,
		GetTempFunction: getTempFunction,
	}
}

// GetTempViaWeb retrieves temperature data via web.
func (ws *WebSensor) GetTempViaWeb() (float32, error) {
	// Check if GetTempFunction is set
	if ws.GetTempFunction == nil {
		return math32.NaN(), errors.New("GetTempFunction is not set")
	}

	// Call the GetTempFunction
	temp, err := ws.GetTempFunction()
	if err != nil {
		return math32.NaN(), fmt.Errorf("error calling GetTempFunction: %v", err)
	}

	return temp, nil
}
