package main

import (
	"fmt"
	"log"
	"url-shortener/config"
	"url-shortener/internal/routes"
	"url-shortener/pkg/database"

	"github.com/spf13/viper"
)

func init() {
	config.LoadDotEnv()

	database.InitializeStore()
}

func main() {
	r := routes.SetupRoutes()

	port := viper.GetString("PORT")

	err := r.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
