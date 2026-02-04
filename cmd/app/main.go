package main

import (
	"log"

	"github.com/evrone/api-servicee/config"
	"github.com/evrone/api-servicee/internal/app"
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
