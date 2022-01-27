package server

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
