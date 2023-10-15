/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

type Validable interface {
	IsValid() bool
}

type StructValidationError struct {
	InvalidData bool
	Errors      []*FieldValidationError
}

type FieldValidationError struct {
	Field      string `json:"field"`
	Validation string `json:"validation"`
	Message    string `json:"message"`
}

func (v *StructValidationError) Error() string {
	if v.InvalidData {
		return "data is invalid or nil"
	}

	return "data validation errors"
}

func (v *StructValidationError) AllErrors() string {
	var sb strings.Builder
	for _, e := range v.Errors {
		sb.Write([]byte(fmt.Sprintf("%s(%s): %s\n", e.Field, e.Validation, e.Message)))
	}
	return sb.String()
}
