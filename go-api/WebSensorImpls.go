package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getTemperatureOpenWeatherMap() (float32, error) {

	// Perform GET request to OpenWeatherMap API
	response, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", 52.127, 6.9133, openWeatherMapAPIKey))
	if err != nil {
		return 0, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response status code: %d", response.StatusCode)
	}

	// Parse the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("error decoding JSON: %v", err)
	}

	// Extract the temperature from the JSON response
	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("unable to parse 'main' data from response")
	}

	temp, ok := mainData["temp"].(float64)
	if !ok {
		return 0, fmt.Errorf("unable to parse 'temp' value from response")
	}

	// Convert temperature from Kelvin to Celsius
	tempCelsius := float32(temp - 273.15)

	return tempCelsius, nil
}
