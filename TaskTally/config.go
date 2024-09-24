package config

import (
	"log"
	"os"
)

var PORT string

func LoadConfig() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
		log.Println("Defaulting to port 8080")
	}
}