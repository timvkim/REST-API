package models

import (
	"errors"
	"fmt"
)

var ErrDoesNotExist = errors.New("does not exist")

type ErrorFields struct {
	Fields map[string]string `json:"fields"`
}

func (e ErrorFields) Error() string {
	return fmt.Sprintf("%v", e.Fields)
}

