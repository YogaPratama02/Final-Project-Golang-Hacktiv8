package validations

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type CustomValidation struct {
	Validator *validator.Validate
}

func DoValidation(i interface{}) map[string]string {
	var message map[string]string

	message = map[string]string{}

	val := CustomValidation{validator.New()}

	if err := val.Validator.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s can not empty!", e.Field())
			case "email":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Input must be email!")
			case "datetime":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s must be date & time", e.Field())
			}
		}

		return message
	}

	return nil
}
