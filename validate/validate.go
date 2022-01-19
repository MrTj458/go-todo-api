package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

// func Struct(data interface{}) ([]*model.ErrorField, bool) {
// 	// Create new validator and have it use json tags for field names
// 	validate := validator.New()
// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 		if name == "-" {
// 			return ""
// 		}
// 		return name
// 	})

// 	// Validate the struct and construct a slice of `ErrorField`s
// 	err := validate.Struct(data)
// 	if err != nil {
// 		var errors []*model.ErrorField
// 		for _, err := range err.(validator.ValidationErrors) {
// 			errors = append(errors, &model.ErrorField{
// 				Location: err.Field(),
// 				Type:     err.Type().String(),
// 				Detail:   err.ActualTag(),
// 			})
// 		}

// 		return errors, false
// 	}

// 	return nil, true
// }
