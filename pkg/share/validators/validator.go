package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetValidator() *validator.Validate {
	return validate
}
func SetUpValidator() {
	validate.RegisterValidation("password", passwordType)
}

func passwordType(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok {
		if len(password) < 8 {
			return false
		}
	}
	return true
}
