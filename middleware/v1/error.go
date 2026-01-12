package v1

import "errors"

var (
	ErrNotExist       = errors.New("value does not exist")
	ErrUnexpectedType = errors.New("unexpected value type for context key")
)
