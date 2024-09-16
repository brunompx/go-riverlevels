package retriever

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/brunompx/go-riverlevels/model"
	"github.com/brunompx/go-riverlevels/repository"
)

// web:
// https://alerta.ina.gob.ar/pub/gui/datosProno&auto=true&timeStart=2024-03-29&timeEnd=2024-07-21&seriesId=3412&siteCode=34&varId=2&calId=289&all=false&format=json
// example:
// https://alerta.ina.gob.ar/pub/datos/datosProno&timeStart=2024-03-29&timeEnd=2024-07-21&seriesId=3412&calId=289&all=false&siteCode=34&varId=2&format=json
const BaseUrlProno = "https://alerta.ina.gob.ar/pub/datos/datosProno"

// web:
// https://alerta.ina.gob.ar/pub/gui/datos&auto=true&timeStart=2023-07-19&timeEnd=2024-07-17&seriesId=34&siteCode=34&varId=2&format=json
// example:
// https://alerta.ina.gob.ar/pub/datos/datos&timeStart=2023-07-19&timeEnd=2024-07-17&seriesId=34&siteCode=34&varId=2&format=json
const BaseUrl = "https://alerta.ina.gob.ar/pub/datos/datos"

func GetData(loc model.Location) model.ForecastResponse {
	parameters := buildParametersMap(loc)
	url := buildUrl(BaseUrlProno, parameters)

	response := GetDataFromAPI(url)

	forecastResponse := unmarshallForecastResponse(response)

	if len(forecastResponse.Data) > 0 {

		//Save data TO file, for manual validation, delete afterwards
		repository.SaveDataAsJsonFile(response, loc)

		return forecastResponse
	}

	responseWeb := GetDataFromWeb(url)

	//Save data TO file, for manual validation, delete afterwards
	repository.SaveDataAsJsonFile(responseWeb, loc)

	return unmarshallForecastResponse(responseWeb)

}

func unmarshallForecastResponse(jsonBytes []byte) model.ForecastResponse {

	//unmarshall json response to ForecastResponse
	var forecastResponse model.ForecastResponse
	err2 := json.Unmarshal(jsonBytes, &forecastResponse)
	if err2 != nil {
		fmt.Println(err2)
	}

	return forecastResponse
}

// NOTE: not working adding params with question mark ?
func buildUrl(baseUrl string, parameters map[string]string) string {

	var url strings.Builder
	url.WriteString(baseUrl)

	for key, value := range parameters {
		url.WriteString("&")
		url.WriteString(key)
		url.WriteString("=")
		url.WriteString(value)
	}
	//fmt.Println(url.String())
	return url.String()
}

func buildParametersMap(loc model.Location) map[string]string {

	timeStart := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	timeEnd := time.Now().AddDate(0, 0, 40).Format("2006-01-02")

	parameters := map[string]string{
		"seriesId":  loc.SeriesId,
		"timeStart": timeStart,
		"timeEnd":   timeEnd,
		"all":       "false",
		"format":    "json",
	}

	if len(loc.VarId) > 0 {
		parameters["varId"] = loc.VarId
	}

	if len(loc.CalId) > 0 {
		parameters["calId"] = loc.CalId
	}

	if len(loc.SiteCode) > 0 {
		parameters["siteCode"] = loc.SiteCode
	}

	if len(loc.CorId) > 0 {
		parameters["corId"] = loc.CorId
	}

	if len(loc.Format) > 0 {
		parameters["format"] = loc.Format
	}

	if len(loc.All) > 0 {
		parameters["all"] = loc.All
	}

	return parameters
}
