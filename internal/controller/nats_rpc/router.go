package v1

import (
	v1 "github.com/evrone/api-servicee/internal/controller/nats_rpc/v1"
	"github.com/evrone/api-servicee/internal/usecase"
	"github.com/evrone/api-servicee/pkg/logger"
	"github.com/evrone/api-servicee/pkg/nats/nats_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewTranslationRoutes(routes, t, l)
	}

	return routes
}
