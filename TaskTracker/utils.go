package utils

import (
    "log"
    "os"
)

func CheckError(err error) {
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
}

func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}