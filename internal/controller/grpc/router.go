package grpc

import (
	v1 "github.com/evrone/api-service/internal/controller/grpc/v1"
	"github.com/evrone/api-service/internal/usecase"
	"github.com/evrone/api-service/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *pbgrpc.Server, t usecase.Translation, l logger.Interface) {
	{
		v1.NewTranslationRoutes(app, t, l)
	}

	reflection.Register(app)
}
