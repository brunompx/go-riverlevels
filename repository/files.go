package repository

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/brunompx/go-riverlevels/model"
)

func SaveDataAsJsonFile(data []byte, loc model.Location) {
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
