package partmap

import "errors"

var (
	ErrEmptyKey   = errors.New("empty key provided")
	ErrEmptyValue = errors.New("empty value provided")
)
