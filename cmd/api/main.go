package main

import (
	"log"
	"url-shortener/internal/routes"
	"url-shortener/pkg/database"
)

func init() {
	database.InitializeStore()
}

func main() {
	r := routes.SetupRoutes()

	err := r.Run(":3000")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
