package v1

import (
	v1 "github.com/evrone/api-service/internal/controller/amqp_rpc/v1"
	"github.com/evrone/api-service/internal/usecase"
	"github.com/evrone/api-service/pkg/logger"
	"github.com/evrone/api-service/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation, l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)

	{
		v1.NewTranslationRoutes(routes, t, l)
	}

	return routes
}
