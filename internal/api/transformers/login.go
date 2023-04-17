package transformers

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ravilock/desafio-backend-lucree/internal/api"
	"github.com/ravilock/desafio-backend-lucree/internal/api/dtos"
	"github.com/ravilock/desafio-backend-lucree/internal/api/validation"
)

func Login(dto *dtos.LoginDto) error {
	*dto.Username = strings.TrimSpace(*dto.Username)
	*dto.Password = strings.TrimSpace(*dto.Password)
	if err := validation.Validate.Struct(dto); err != nil {
		if validationErrors := new(validator.ValidationErrors); errors.As(err, validationErrors) {
			for _, validationError := range *validationErrors {
				return api.InvalidFieldError(validationError.Field(), validationError.Value(), "")
			}
		}
	}

	return nil
}
