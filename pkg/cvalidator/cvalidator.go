package cvalidator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
	once      sync.Once
)

const (
	ErrorValidator = "ERROR_VALIDATOR"
)

func nonEmptyString(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return str != ""
}

func init() {
	once.Do(func() {
		Validator = validator.New()

		Validator.RegisterValidation("nonempty", nonEmptyString)
	})
}
