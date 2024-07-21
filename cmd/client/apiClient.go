package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/brunompx/go-riverlevels/cmd/model"
)

const BaseUrl = "https://alerta.ina.gob.ar/pub/datos/datosProno"

func GetData(loc model.Location) []byte {
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
	//fmt.Println(string(responseBody))

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
