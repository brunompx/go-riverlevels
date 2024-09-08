package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/brunompx/go-riverlevels/client"
	"github.com/brunompx/go-riverlevels/db"
	"github.com/brunompx/go-riverlevels/handlers"
	"github.com/brunompx/go-riverlevels/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "postgres", "riverlevels", "5432")
	db, err := newDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	//db.AutoMigrate(&model.Forecast{})
	//db.AutoMigrate(&model.ForecastSet{})
	//db.AutoMigrate(&model.ForecastLevel{})

	now := time.Now()
	fmt.Println("locations cargadas")
	//retrieveData()
	fmt.Println("Tardo en totral: ", time.Since(now))

	processRosarioFile()
	pruebadechat()

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HandleHome)

	router.HandleFunc("GET /linechart", handlers.HandleLineChart)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func retrieveData() {
	locations := getLocationsData()

	var wg sync.WaitGroup
	wg.Add(len(locations.Locations))
	for _, location := range locations.Locations {
		loc := location
		go processLocation(loc, &wg)
	}
	wg.Wait()
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

func newDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func initStorage(db *gorm.DB) {
	genericDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	pingErr := genericDB.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")
}

// just testing unmarshall to struct /////////////////////////////////////////////////////////////////////////////////////////////////
func processRosarioFile() {

	jsonFile, err := os.Open("rosario1.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened locations.json")
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var forecastResponse model.ForecastResponse
	err2 := json.Unmarshal(byteValue, &forecastResponse)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%+v\n", forecastResponse)
	fmt.Println("-----------------------------------------------")
	for _, level := range forecastResponse.Data {
		fmt.Printf("%+v\n", level)
		fmt.Println("-----------------------------------------------")
	}

}

func pruebadechat() {
	// Your JSON data
	jsonData := `
{
	"responseHeader": {
		"varid": 2,
		"request": "datosProno",
		"estacion_tabla": "alturas_prefe",
		"corid": 857049,
		"timeEnd": "20241015T000000",
		"seriesid": 3412,
		"proc_nombre": "pronóstico",
		"unitid": 11,
		"sitecode": 34,
		"estacion_nombre": "Rosario",
		"red_nombre": "escalas Prefectura Nacional",
		"calid": 289,
		"cal_name": "tabprono_central",
		"responseTimestamp": "2024-09-05T18:25:00",
		"timeStart": "20240826T000000",
		"var_nombre": "Altura hidrométrica",
		"cal_model": "tabprono",
		"queryUrl": "request=datosProno&seriesId=3412&timeStart=2024-08-26&timeEnd=2024-10-15&varId=2&calId=289&corId=857049&all=false&format=json",
		"model_id": 60,
		"creationTime": "2024-09-05T18:25:00-03:00",
		"unit_nombre": "metros",
		"forecastdate": "2024-09-03T00:00:00"
	},
	"data": [
		{
			"prono_id": 904704811,
			"timestart": "2024-09-03T00:00:00",
			"timeend": "2024-09-03T00:00:00",
			"valor": 1.09
		},
		{
			"prono_id": 904704805,
			"timestart": "2024-09-03T00:00:00",
			"timeend": "2024-09-03T00:00:00",
			"valor": 1.09
		}
	]
}
`

	// Create a variable of type ForecastResponse to hold the unmarshaled data
	var forecast model.ForecastResponse

	// Custom time format to parse the timestamp fields
	const timeLayout = "2006-01-02T15:04:05"

	// Unmarshal the JSON into the forecast struct
	err := json.Unmarshal([]byte(jsonData), &forecast)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Manually parse time fields if necessary
	// (json.Unmarshal won't automatically handle custom time formats)
	for i, entry := range forecast.Data {
		start, err := time.Parse(timeLayout, entry.TimeStart.Format(timeLayout))
		if err != nil {
			fmt.Println("Error parsing TimeStart:", err)
		}
		end, err := time.Parse(timeLayout, entry.TimeEnd.Format(timeLayout))
		if err != nil {
			fmt.Println("Error parsing TimeEnd:", err)
		}
		forecast.Data[i].TimeStart = start
		forecast.Data[i].TimeEnd = end
	}

	// Print the result
	fmt.Println("#########################################################")
	fmt.Printf("%+v\n", forecast)
	fmt.Println("-----------------------------------------------")
	for _, level := range forecast.Data {
		fmt.Printf("%+v\n", level)
		fmt.Println("-----------------------------------------------")
	}
}
