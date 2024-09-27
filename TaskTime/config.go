package config

import (
	"encoding/json"
	"os"
	"./utils"
)

type Config struct {
	ServerPort string `json:"server_port"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		utils.LogError("Failed to open config file")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		utils.LogError("Failed to decode config file")
		return
	}

	utils.LogInfo("Config loaded successfully")
}