package v1

import (
	v1 "github.com/evrone/api-servicee/docs/proto/v1"
	"github.com/evrone/api-servicee/internal/usecase"
	"github.com/evrone/api-servicee/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.TranslationServer

	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}
