package utils

import (
	"ecommerce-v1/auth/models"
	"strings"

	"github.com/go-playground/validator"
)

func StructValidator(model interface{}) []*models.StructError {
	var errors []*models.StructError
	var err error

	validate := validator.New()
	err = validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &models.StructError{Field: strings.ToLower(err.Field()), Tag: err.ActualTag(), Value: err.Param()})
		}
		return errors
	}
	return nil
}
