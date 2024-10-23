package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	EnvType       string `mapstructure:"ENV_TYPE"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	WebappBaseUrl string `mapstructure:"WEBAPP_BASE_URL"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + ".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
	}
	return
}
