package clickmeeting

import (
	"fmt"
	"strings"
)

type APIError struct {
	Code   int    `json:"code"`
	Name   string `json:"name"`
	Errors []struct {
		Name    interface{} `json:"name"`
		Message string      `json:"message"`
	} `json:"errors"`
}

func (e APIError) Error() string {
	errors := make([]string, 0, len(e.Errors))
	for _, err := range e.Errors {
		errors = append(errors, err.Message)
	}

	return fmt.Sprintf("%d-%s: %s", e.Code, e.Name, strings.Join(errors, ","))
}

func (e APIError) Is(target error) bool {
	_, ok := target.(APIError)
	return ok
}
