package main

import (
	"net/http"

	"github.com/brunompx/go-riverlevels/database"
	"github.com/brunompx/go-riverlevels/handlers"
	"github.com/brunompx/go-riverlevels/ingestor"
	"github.com/brunompx/go-riverlevels/repository"
	"github.com/brunompx/go-riverlevels/service"
)

func main() {

	db := database.GetDatabase()
	repositories := repository.InitRepositories(db)
	services := service.InitServices(repositories)

	ingestor.IngestData(services)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HandleHome)

	router.HandleFunc("GET /linechart", handlers.HandleLineChart)

	//server := http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//server.ListenAndServe()
}
