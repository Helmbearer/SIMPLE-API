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

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
