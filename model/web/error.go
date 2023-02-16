package web

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Message string      `json:"-"`
	Code    string      `json:"code"`
	Details interface{} `json:"details"`
}

func (v ResponseError) Error() string {
	return v.Message
}

func NewResponseError(
	message string,
	code string,
) ResponseError {
	return ResponseError{
		Message: message,
		Code:    code,
		Details: []string{},
	}
}

func NewResponseErrorFromError(
	err error,
	code string,
) ResponseError {
	return NewResponseError(err.Error(), code)
}

func NewResponseErrorFromValidator(
	err validator.ValidationErrors,
) ResponseError {
	errList := []map[string]string{}
	for _, err := range err {
		errList = append(errList, map[string]string{
			"field": err.Field(),
			"tag":   err.Tag(),
		})
	}

	return ResponseError{
		Message: "input is not valid",
		Code:    InvalidInput,
		Details: errList,
	}
}

func NewResponseErrorf(
	code string,
	format string,
	value ...interface{},
) ResponseError {
	return NewResponseErrorFromError(
		fmt.Errorf(format, value...),
		code,
	)
}
