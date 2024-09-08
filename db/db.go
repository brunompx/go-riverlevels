package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brunompx/go-riverlevels/model"
)

func SaveData(data []byte, loc model.Location) {
	fileName := "data-" + loc.SeriesId + "-" + loc.SiteCode + "-" + loc.CorId + "-" + loc.CalId + ".json"
	var prettyJSON bytes.Buffer
	if len(data) > 0 {
		error := json.Indent(&prettyJSON, data, "", "\t")
		if error != nil {
			log.Fatal(error)
		}

		toStructs(prettyJSON.Bytes())

		err := os.WriteFile(fileName, prettyJSON.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func toStructs(jsonBytes []byte) {

	var forecastResponse model.ForecastResponse

	// Unmarshal the JSON into the person struct
	err := json.Unmarshal(jsonBytes, &forecastResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the result
	fmt.Printf("%+v\n", forecastResponse)
}
