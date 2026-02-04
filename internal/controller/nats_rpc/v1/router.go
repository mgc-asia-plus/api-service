package v1

import (
	"github.com/evrone/api-service/internal/usecase"
	"github.com/evrone/api-service/pkg/logger"
	"github.com/evrone/api-service/pkg/nats/nats_rpc/server"
	"github.com/go-playground/validator/v10"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(routes map[string]server.CallHandler, t usecase.Translation, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		routes["v1.getHistory"] = r.getHistory()
	}
}
