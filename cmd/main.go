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
	"github.com/brunompx/go-riverlevels/db"
	"github.com/brunompx/go-riverlevels/handlers"
	"github.com/brunompx/go-riverlevels/model"
)

func main() {

	now := time.Now()

	locations := getLocationsData()

	fmt.Println("locations cargadas")

	var wg sync.WaitGroup
	wg.Add(len(locations.Locations))

	for _, location := range locations.Locations {
		loc := location
		go processLocation(loc, &wg)
	}
	wg.Wait()

	fmt.Println("Tardo en totral: ", time.Since(now))

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HandleHome)

	router.HandleFunc("GET /linechart", handlers.HandleLineChart)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func processLocation(loc model.Location, wg *sync.WaitGroup) {

	defer wg.Done()

	responseData := client.GetData(loc)
	db.SaveData(responseData, loc)
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
