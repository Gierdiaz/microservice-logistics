package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validator(s interface{}) error {
	return validate.Struct(s)
}
