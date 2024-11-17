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

const StTypeForecast = "forecast"
const StTypeMeasure = "measure"

func IngestData(services *service.Service) {

	//TODO save data from file, remove this test when we can get consistent data from alerta.ina.gob.ar
	//fore := repository.ProcessSavedFileFile()
	//services.ForecastService.Save(&fore)
	// fin TODO esto es un test borrar despues

	stations := getStationsData()

	now := time.Now()
	fmt.Println("stations cargadas")

	var wg sync.WaitGroup
	wg.Add(len(stations.Stations))
	for _, station := range stations.Stations {
		st := station
		go processLocation(st, services, &wg)
	}
	wg.Wait()

	fmt.Println("Tardo en totral: ", time.Since(now))
}

func processLocation(st types.Station, services *service.Service, wg *sync.WaitGroup) {

	defer wg.Done()

	if st.StationType == StTypeForecast {
		processForecastLoc(st, services)
	} else if st.StationType == StTypeMeasure {
		processMeasureLoc(st, services)
	}
}

func processForecastLoc(st types.Station, services *service.Service) {
	forecastResponse := retriever.GetForecastData(st)
	forecast := forecastResponse.NormalizeToForecast()
	services.ForecastService.Save(&forecast)
}

func processMeasureLoc(st types.Station, services *service.Service) {
	measureResponse := retriever.GetMeasuretData(st)
	measure := measureResponse.NormalizeToMeasure()
	services.MeasureService.Save(&measure)
}

func getStationsData() types.Stations {
	jsonFile, err := os.Open("stations.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened locations.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var stations types.Stations
	json.Unmarshal(byteValue, &stations)
	return stations
}
