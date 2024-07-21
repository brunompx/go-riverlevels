package db

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func SaveData(data []byte, seriesId string, siteCode string, calId string) {

	fileName := "data-" + seriesId + "-" + siteCode + "-" + calId + ".json"

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
