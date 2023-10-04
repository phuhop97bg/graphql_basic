package utils

import "backend-food/pkg/share/validators"

func CheckValidate(x interface{}) error {
	return validators.GetValidator().Struct(x)
}
