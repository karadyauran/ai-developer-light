package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	EnvType      string `mapstructure:"ENV_TYPE"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	GrpcAuthPort string `mapstructure:"GPRC_AUTH_PORT"`

	GitHubClientID     string `mapstructure:"GITHUB_CLIENT_ID"`
	GitHubClientSecret string `mapstructure:"GITHUB_CLIENT_SECRET"`
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
