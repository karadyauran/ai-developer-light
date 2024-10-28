package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	EnvType             string `mapstructure:"ENV_TYPE"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	GrpcAIGeneratorPort string `mapstructure:"GPRC_API_GENERATOR_PORT"`

	OpenAIAPI string `mapstructure:"OPENAI_API_KEY"`
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
