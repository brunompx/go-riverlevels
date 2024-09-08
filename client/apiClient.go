package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/brunompx/go-riverlevels/model"
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

func GetData(loc model.Location) []byte {

	parameters := buildParametersMap(loc)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, buildUrl(BaseUrlProno, parameters), nil)
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

	//fmt.Println(resp.Status)
	//fmt.Println(string(responseBody))

	return responseBody
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

	return parameters
}
