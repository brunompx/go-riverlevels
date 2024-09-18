package main

import (
	"net/http"

	"github.com/brunompx/go-riverlevels/handlers"
)

func main() {

	//db := database.GetDatabase()
	//repositories := repository.InitRepositories(db)
	//services := service.InitServices(repositories)

	//ingestor.IngestData(services)

	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.HandleHome)

	router.HandleFunc("GET /linechart", handlers.HandleLineChart)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
