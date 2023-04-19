package utils

import "errors"

// Predefined errors
var (
	ErrInvalidReader   = errors.New("invalid reader")
	ErrEmptyInput      = errors.New("empty input")
	ErrInvalidPartSize = errors.New("part size must be greater than 0")
)
