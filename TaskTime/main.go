package main

import (
	"log"
	"net/http"
	"./routes"
	"./config"
)

func main() {
	config.LoadConfig()

	r := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}