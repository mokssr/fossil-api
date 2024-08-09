package api

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

var validationMessageFormat = map[string]string{
	"required": "Field %s is required",
	"email":    "Field %s must be a valid email address",
}

func BootstrapRequestValidator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})
}

// Returns mapped error message field => [msg1, msg2]
func ValidateStruct(src interface{}) (bool, map[string][]string) {

	var errorBags = map[string][]string{}

	if err := Validate.Struct(src); err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
			msg := humanizeMessage(e)
			errorBags[e.Field()] = append(errorBags[e.Field()], msg)
		}

		return false, errorBags
	}

	return true, nil
}

func humanizeMessage(err validator.FieldError) string {
	return fmt.Sprintf(validationMessageFormat[err.Tag()], err.Field())
}
