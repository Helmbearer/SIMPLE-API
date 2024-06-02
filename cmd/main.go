package main

import (
	"SIMPLE-API/config"
	"SIMPLE-API/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	router := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
