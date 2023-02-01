package models

import "fmt"

type ErrorFields struct {
	Fields map[string]string `json:"fields"`
}

func (e ErrorFields) Error() string {
	return fmt.Sprintf("%v", e.Fields)
}
