package partmap

import "errors"

var (
	// ErrEmptyKey is returned when empty key is provided.
	ErrEmptyKey = errors.New("empty key provided")
	// ErrNotFound is returned when key is not found in map.
	ErrNotFound = errors.New("key not found")
)
