package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		PG      PG
		GRPC    GRPC
		RMQ     RMQ
		NATS    NATS
		Metrics Metrics
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT"` // Falls back to PORT if not set
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
		AuthToken      string `env:"HTTP_AUTH_TOKEN" envDefault:""` // if set, require X-Auth-Token header
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	// GRPC -.
	GRPC struct {
		Port string `env:"GRPC_PORT"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER"`
		ClientExchange string `env:"RMQ_RPC_CLIENT"`
		URL            string `env:"RMQ_URL"`
	}

	// NATS -.
	NATS struct {
		ServerExchange string `env:"NATS_RPC_SERVER"`
		URL            string `env:"NATS_URL"`
	}

	// Metrics -.
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	// Swagger -.
	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// Railway compatibility: Use PORT if HTTP_PORT is not set
	if cfg.HTTP.Port == "" {
		if port := os.Getenv("PORT"); port != "" {
			cfg.HTTP.Port = port
		} else {
			return nil, fmt.Errorf("config error: HTTP_PORT or PORT environment variable is required")
		}
	}

	return cfg, nil
}
