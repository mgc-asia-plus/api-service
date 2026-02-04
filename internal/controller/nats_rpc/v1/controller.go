package v1

import (
	"github.com/evrone/api-service/internal/usecase"
	"github.com/evrone/api-service/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
