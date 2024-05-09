package main

import (
	"encoding/json"
	"fmt"
	"github.com/rkusa/gm/math32"
	"net/http"
	"time"
)

const (
	openWeatherMapURL     = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s"
	esp32SensorURL        = "http://esp32-temperatur-sensor.fritz.box/"
	WebSensorPrefix       = "WEB_SENSOR_" //Every outside sensor has to use the SENSOR_PREFIX
	SensorOpenWeatherName = WebSensorPrefix + "OPEN_WEATHER_API"
	SensorESP32Name       = WebSensorPrefix + "ESP32"
)

var WebSensors []WebSensor

func initWebSensors() {
	if !UseWebSensors {
		Log.Println("do not use Web sensors")
		return
	}

	WebSensors = make([]WebSensor, 0, 2)

	openWeatherMapSensor := NewWebSensor(SensorOpenWeatherName, getTemperatureOpenWeatherMap)
	WebSensors = append(WebSensors, *openWeatherMapSensor)

	openWeatherMapSensor = NewWebSensor(SensorESP32Name, getTemperatureESP32)
	WebSensors = append(WebSensors, *openWeatherMapSensor)
}

func getAllWebSensorIDs() []string {
	sensorNames := make([]string, 0, len(WebSensors))

	for _, webSensor := range WebSensors {
		sensorNames = append(sensorNames, webSensor.SensorID)
	}

	return sensorNames
}

func getWebSensor(sensorID string) *WebSensor {
	for _, webSensor := range WebSensors {
		if webSensor.SensorID == sensorID {
			return &webSensor
		}
	}
	return nil
}

// getTemperatureOpenWeatherMap retrieves the current temperature from OpenWeatherMap.
func getTemperatureOpenWeatherMap() (float32, error) {
	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// Perform GET request to OpenWeatherMap API
	response, err := client.Get(fmt.Sprintf(openWeatherMapURL, 52.127, 6.9133, API_KEY))
	if err != nil {
		return math32.NaN(), fmt.Errorf("error making HTTP request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK
	if response.StatusCode != http.StatusOK {
		return math32.NaN(), fmt.Errorf("unexpected response status code: %d", response.StatusCode)
	}

	// Parse the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return math32.NaN(), fmt.Errorf("error decoding JSON: %v", err)
	}

	// Extract the temperature from the JSON response
	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		return math32.NaN(), fmt.Errorf("unable to parse 'main' data from response")
	}

	temp, ok := mainData["temp"].(float64)
	if !ok {
		return math32.NaN(), fmt.Errorf("unable to parse 'temp' value from response")
	}

	// Convert temperature from Kelvin to Celsius
	tempCelsius := float32(temp - 273.15)

	return tempCelsius, nil
}

// getTemperatureESP32 retrieves the temperature from an ESP32 temperature sensor.
func getTemperatureESP32() (float32, error) {
	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	// Perform GET request to the ESP32 temperature sensor
	response, err := client.Get(esp32SensorURL)
	if err != nil {
		return math32.NaN(), fmt.Errorf("error making HTTP request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK
	if response.StatusCode != http.StatusOK {
		return math32.NaN(), fmt.Errorf("unexpected response status code: %d", response.StatusCode)
	}

	// Parse the JSON response
	var data map[string]float32
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return math32.NaN(), fmt.Errorf("error decoding JSON: %v", err)
	}

	// Extract the temperature from the JSON response
	temperature, ok := data["temperature"]
	if !ok {
		return math32.NaN(), fmt.Errorf("unable to parse 'temperature' value from response")
	}
	if temperature <= -127+0.001 {
		return math32.NaN(), fmt.Errorf("sensor issue temperature is below -127 degrees Celsisus")
	}

	return temperature, nil
}
