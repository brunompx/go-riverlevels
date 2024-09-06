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

	now := time.Now()
	fmt.Println("locations cargadas")
	//retrieveData()
	fmt.Println("Tardo en totral: ", time.Since(now))

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
