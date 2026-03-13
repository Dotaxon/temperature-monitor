package ds18b20wrapper

import (
	"fmt"
	"go-api/logger"
	"math"
	"sync"
	"time"

	"github.com/yryz/ds18b20"
)

var temps = make(map[string]float32)
var tempsMutex sync.RWMutex

func init() {
	refresh()
	go timer()
}

func timer() {
	for {
		refresh()
		time.Sleep(5 * time.Second)
	}
}

func refresh() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		logger.Log.Println(err)
		return
	}

	newTemps := make(map[string]float32, len(sensors))

	for _, sensor := range sensors {
		temp, err := ds18b20.Temperature(sensor)
		if err != nil {
			temp = math.NaN()
		}
		newTemps[sensor] = float32(temp)
	}

	tempsMutex.Lock()
	defer tempsMutex.Unlock()
	temps = newTemps
}

func GetSensors() ([]string, error) {
	tempsMutex.RLock()
	defer tempsMutex.RUnlock()
	sensors := make([]string, len(temps))
	i := 0
	for sensor := range temps {
		sensors[i] = sensor
		i++
	}
	return sensors, nil
}

func GetTemp(name string) (float32, error) {
	tempsMutex.RLock()
	defer tempsMutex.RUnlock()
	temp, ok := temps[name]
	if !ok {
		return float32(math.NaN()), fmt.Errorf("sensor %s not found", name)
	}
	return temp, nil
}
