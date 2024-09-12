package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/brunompx/go-riverlevels/client"
	"github.com/brunompx/go-riverlevels/database"
	"github.com/brunompx/go-riverlevels/handlers"
	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/storage"
)

func main() {

	db := database.GetDatabase()

	now := time.Now()
	fmt.Println("locations cargadas")
	//retrieveData()
	fmt.Println("Tardo en totral: ", time.Since(now))

	processRosarioFile()

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HandleHome)

	router.HandleFunc("GET /linechart", handlers.HandleLineChart)

	//server := http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//server.ListenAndServe()
}

func retrieveData() {
	locations := getLocationsData()

	var wg sync.WaitGroup
	wg.Add(len(locations.Locations))
	for _, location := range locations.Locations {
		loc := location
		go processLocation(loc, &wg)
	}
	wg.Wait()
}

func processLocation(loc model.Location, wg *sync.WaitGroup) {

	defer wg.Done()

	responseData := client.GetData(loc)
	storage.SaveData(responseData, loc)
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

// just testing unmarshall to struct /////////////////////////////////////////////////////////////////////////////////////////////////
func processRosarioFile() {
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	fmt.Println("&&&&&&&&&&&&&&& processRosarioFile:")

	jsonFile, err := os.Open("rosario1.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened locations.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var forecastResponse model.ForecastResponse
	err2 := json.Unmarshal(byteValue, &forecastResponse)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", forecastResponse)
	fmt.Println("-----------------------------------------------")
	for _, level := range forecastResponse.Data {
		fmt.Printf("%+v\n", level.PronoId)
		fmt.Printf("%+v\n", level.TimeStart)
		fmt.Printf("%+v\n", level.Valor)
		fmt.Println("-----------------------------------------------")
	}

	fmt.Println("-#########################################################################################-")

	normalized := forecastResponse.NormalizeToForecast()

	fmt.Printf("%+v\n", normalized.ForecastDate)

	for _, set := range normalized.ForecastSets {
		fmt.Println("-##################-")
		for _, level := range set.ForecastLevels {
			fmt.Printf("%+v\n", level.PronoID)
			fmt.Printf("%+v\n", level.TimeProno)
			fmt.Printf("%+v\n", level.Valor)
			fmt.Println("-----------------------------------------------")
		}

	}

}
