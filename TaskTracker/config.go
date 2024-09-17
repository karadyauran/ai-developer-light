package config

import (
    "project/utils"
    "os"
)

type Config struct {
    ServerAddress string
}

func LoadConfig() (*Config, error) {
    serverAddress := utils.GetEnv("SERVER_ADDRESS", ":8080")
    return &Config{
        ServerAddress: serverAddress,
    }, nil
}

func init() {
    os.Setenv("SERVER_ADDRESS", ":8080")
}