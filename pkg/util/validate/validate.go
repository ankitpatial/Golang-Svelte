/*
	Simsaw Software Pvt. Ltd.
	Author: Ankit Patial <ankit@simsaw.com>
*/

package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var (
	DefaultErrMsg    = "'%s' validation '%s' failed."
	RequiredErrMsg   = "required"
	EmailErrMsg      = "'%s' is not a valid email address"
	ValidErrMsg      = "value '%s' is not allowed"
	ValidEmptyErrMsg = "value must not be empty"
	OneOfErrMsg      = "value must be one of '%s'."
	UrlErrMsg        = "invalid url"
)

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	_ = validate.RegisterValidation("valid", isValid)
}

func isValid(fl validator.FieldLevel) bool {
	if v, ok := fl.Field().Interface().(Validable); ok {
		return v.IsValid()
	}

	return true
}

func IsStructValidationErr(err error) bool {
	switch err.(type) {
	case *StructValidationError:
		return true
	default:
		return false
	}
}

func Struct(s interface{}) *StructValidationError {
	err := validate.Struct(s)
	if err == nil { // no error
		return nil
	}

	valErr := new(StructValidationError)
	if _, ok := err.(*validator.InvalidValidationError); ok {
		// struct is not valid
		valErr.InvalidData = true
		valErr.Errors = append(valErr.Errors, &FieldValidationError{
			Field:      "",
			Validation: "valid",
			Message:    "invalid struct",
		})
		return valErr
	}

	// field validation issues
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		ns := v.Namespace()
		idx := strings.Index(ns, ".") + 1
		field := ns[idx:]

		fieldErr := &FieldValidationError{
			Field:      field,
			Validation: v.Tag(),
		}

		switch v.Tag() {
		case "required", "required_unless", "required_if":
			fieldErr.Message = RequiredErrMsg
		case "email":
			fieldErr.Message = fmt.Sprintf(EmailErrMsg, v.Value())
		case "valid":
			val := fmt.Sprintf("%v", v.Value())
			if val == "" {
				fieldErr.Message = ValidEmptyErrMsg
			} else {
				fieldErr.Message = fmt.Sprintf(ValidErrMsg, v.Value())
			}

		case "oneof":
			fieldErr.Message = fmt.Sprintf(OneOfErrMsg, v.Param())
		case "url":
			fieldErr.Message = UrlErrMsg

		default:
			fieldErr.Message = fmt.Sprintf(DefaultErrMsg, field, v.Tag())
		}

		valErr.Errors = append(valErr.Errors, fieldErr)
	}

	return valErr
}

// Email address validation will ensure none empty validate email address
func Email(email string) bool {
	if err := validate.Var(email, "required,email"); err != nil {
		return false
	}

	return true
}
