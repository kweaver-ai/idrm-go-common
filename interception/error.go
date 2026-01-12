package interception

import "errors"

var (
	ErrNotExist       = errors.New("token required context token does not exist")
	ErrUnexpectedType = errors.New("unexpected value type for token context key")
)
