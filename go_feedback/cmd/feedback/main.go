package main

import (
	"log"
	"ms-feedback/docs"
	"ms-feedback/internal/config"
	"ms-feedback/internal/server"
)

// @title Basic go API for ms-feedback
// @version 1.0
// @description help to basic app
// @BasePath /ms-feedback
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name ms-feedback
// @schemes https http
func main() {
	log.Println("ms-feedback version 1.0.1")

	cfg := config.MustLoadConfig()

	docs.SwaggerInfo.Host = cfg.SwaggerHost

	if err := server.Start(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
