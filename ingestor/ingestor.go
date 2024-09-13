package ingestor

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
	"github.com/brunompx/go-riverlevels/service"
)

func IngestData(services *service.Service) {
	locations := getLocationsData()

	now := time.Now()
	fmt.Println("locations cargadas")

	var wg sync.WaitGroup
	wg.Add(len(locations.Locations))
	for _, location := range locations.Locations {
		loc := location
		go processLocation(loc, services, &wg)
	}
	wg.Wait()

	fmt.Println("Tardo en totral: ", time.Since(now))
}

func processLocation(loc model.Location, services *service.Service, wg *sync.WaitGroup) {

	defer wg.Done()

	responseData := GetData(loc)
	forecast := responseToForecast(responseData)
	services.ForecastService.Save(&forecast)
	repository.SaveDataAsJsonFile(responseData, loc)
}

func responseToForecast(jsonBytes []byte) model.Forecast {

	//unmarshall json response to ForecastResponse
	var forecastResponse model.ForecastResponse
	err2 := json.Unmarshal(jsonBytes, &forecastResponse)
	if err2 != nil {
		fmt.Println(err2)
	}

	//transform to Forecast struct
	normalized := forecastResponse.NormalizeToForecast()

	return normalized

}

func getLocationsData() model.Locations {
	jsonFile, err := os.Open("locations.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var locations model.Locations
	json.Unmarshal(byteValue, &locations)
	return locations
}
