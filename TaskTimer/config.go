package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port string `json:"port"`
}

var Config Configuration

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Could not open config file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("Could not decode config JSON", err)
	}
}