package configs

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfig *configs

type configs struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

// InitEnvConfigs initializes the env configs
func InitEnvConfigs() {
	EnvConfig = loadEnvVariables()
}

// LoadEnvVariables loads the env variables from the .env file
func loadEnvVariables() (config *configs) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}
	return
}
