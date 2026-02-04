package main

import (
	"log"

	"github.com/evrone/api-service/config"
	"github.com/evrone/api-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
