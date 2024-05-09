package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Temperature float32 `json:"temperature"`
}

func GetRequestUrl_OPENWEATHER() string {
	if API_KEY == "toChange" || LONGITUDE == "toChange" || LATITUDE == "toChange" {
		Log.Println("Missing part in Request Url please change ApiKey.go")
		return ""
	}
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", API_KEY, LATITUDE, LONGITUDE)
}

func GetRequestUrl_ESP32() string {
	return "http://esp32-temperatur-sensor.fritz.box/"
}

func GetTempViaWeb() (float32, error) {
	requestUrl := GetRequestUrl_ESP32()

	client := &http.Client{
		Timeout: 2 * time.Second, // Timeout set to 1 second
	}

	response, err := client.Get(requestUrl)
	if err != nil {
		return 0, fmt.Errorf("could not get temperature from api error: %s", err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		Log.Println("Error reading response body:", err)
		return 0, fmt.Errorf("could not get temperature from api error: %s", err.Error())
	}

	var parsedResponse Response
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		Log.Println("Error reading response body:", err)
		return 0, fmt.Errorf("could not get temperature from api error: %s", err.Error())
	}

	return parsedResponse.Temperature, nil
}
