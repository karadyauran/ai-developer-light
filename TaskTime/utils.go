package utils

import (
	"log"
	"os"
)

func LogInfo(message string) {
	log.SetOutput(os.Stdout)
	log.Println("[INFO]: " + message)
}

func LogError(message string) {
	log.SetOutput(os.Stderr)
	log.Println("[ERROR]: " + message)
}

func CheckErr(err error, context string) {
	if err != nil {
		LogError(context + ": " + err.Error())
	}
}