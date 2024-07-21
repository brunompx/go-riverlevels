package db

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/brunompx/go-riverlevels/cmd/model"
)

func SaveData(data []byte, loc model.Location) {

	fileName := "data-" + loc.SeriesId + "-" + loc.SiteCode + "-" + loc.CalId + ".json"

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, data, "", "\t")
	if error != nil {
		log.Fatal(error)
	}

	err := os.WriteFile(fileName, prettyJSON.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

}