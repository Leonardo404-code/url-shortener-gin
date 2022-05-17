package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

const CacheDuration = 6 * time.Hour

func LoadDotEnv() {
	viper.AddConfigPath(".")

	viper.SetConfigName(".env.local")

	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("error reading env file, %v", err)
	}
}
