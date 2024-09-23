package ingestor

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/brunompx/go-riverlevels/ingestor/retriever"
	"github.com/brunompx/go-riverlevels/types"

	"github.com/brunompx/go-riverlevels/service"
)

func IngestData(services *service.Service) {

	//TODO save data from file, remove this test when we can get consistent data from alerta.ina.gob.ar
	//fore := repository.ProcessSavedFileFile()
	//services.ForecastService.Save(&fore)
	// fin TODO esto es un test borrar despues

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

func processLocation(loc types.Location, services *service.Service, wg *sync.WaitGroup) {

	defer wg.Done()

	forecastResponse := retriever.GetData(loc)
	forecast := forecastResponse.NormalizeToForecast()
	services.ForecastService.Save(&forecast)
}

func getLocationsData() types.Locations {
	jsonFile, err := os.Open("locations.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var locations types.Locations
	json.Unmarshal(byteValue, &locations)
	return locations
}
