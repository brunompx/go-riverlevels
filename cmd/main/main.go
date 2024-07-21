package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/brunompx/go-riverlevels/cmd/client"
	"github.com/brunompx/go-riverlevels/cmd/db"
	"github.com/brunompx/go-riverlevels/cmd/model"
)

func main() {

	now := time.Now()

	locations := getLocationsData()

	var wg sync.WaitGroup
	wg.Add(len(locations.Locations))

	for _, location := range locations.Locations {
		loc := location
		go processLocation(loc, &wg)
	}
	wg.Wait()

	fmt.Println("Tardo en totral: ", time.Since(now))
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
