package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	EnvType       string `mapstructure:"ENV_TYPE"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	OpenApiKey    string `mapstructure:"OPENAI_API_KEY"`
	WebappBaseUrl string `mapstructure:"WEBAPP_BASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + ".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	return
}
