package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brunompx/go-riverlevels/cmd/db"
)

type Location struct {
	SeriesId string `json:"seriesId"`
	SiteCode string `json:"siteCode"`
	CalId    string `json:"calId"`
}

type Locations struct {
	Locations []Location `json:"locations"`
}

const BaseUrl = "https://alerta.ina.gob.ar/pub/datos/datosProno"

func main() {

	now := time.Now()

	locations := getLocationsData()

	for _, location := range locations.Locations {
		processLocation(location)
	}
	fmt.Println("Tardo en totral: ", time.Since(now))
}

func processLocation(loc Location) {

	responseData := getData(loc)
	db.SaveData(responseData, loc.SeriesId, loc.SiteCode, loc.CalId)

}

func getData(loc Location) []byte {
	parameters := buildParametersMap(loc.SeriesId, loc.SiteCode, loc.CalId)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, buildUrl(BaseUrl, parameters), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))

	return responseBody

}

// TODO: see why is not working with question mark ?
func buildUrl(baseUrl string, parameters map[string]string) string {
	var url strings.Builder
	url.WriteString(baseUrl)

	for key, value := range parameters {
		url.WriteString("&")
		url.WriteString(key)
		url.WriteString("=")
		url.WriteString(value)
	}
	return url.String()
}

func buildParametersMap(seriesId string, siteCode string, calId string) map[string]string {

	timeStart := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	timeEnd := time.Now().AddDate(0, 0, 40).Format("2006-01-02")

	parameters := map[string]string{
		"seriesId":  seriesId,
		"siteCode":  siteCode,
		"calId":     calId,
		"timeStart": timeStart,
		"timeEnd":   timeEnd,
		"varId":     "2",
		"all":       "false",
		"format":    "json",
	}

	return parameters
}

func getLocationsData() Locations {

	jsonFile, err := os.Open("locations.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var locations Locations
	json.Unmarshal(byteValue, &locations)
	return locations
}
