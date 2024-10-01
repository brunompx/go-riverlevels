package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/brunompx/go-riverlevels/types"
)

func SaveDataAsJsonFile(data []byte, loc types.Station) {
	fileName := "data-" + loc.SeriesId + "-" + loc.SiteCode + "-" + loc.CorId + "-" + loc.CalId + ".json"
	var prettyJSON bytes.Buffer
	if len(data) > 0 {
		error := json.Indent(&prettyJSON, data, "", "\t")
		if error != nil {
			log.Fatal(error)
		}
		err := os.WriteFile(fileName, prettyJSON.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ProcessSavedFileFile() types.Forecast {

	jsonFile, err := os.Open("rosario2.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var forecastResponse types.ForecastResponse
	err2 := json.Unmarshal(byteValue, &forecastResponse)
	if err2 != nil {
		fmt.Println(err2)
	}

	normalized := forecastResponse.NormalizeToForecast()

	return normalized
}
