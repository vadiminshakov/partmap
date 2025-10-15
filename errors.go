package partmap

import "errors"

var (
	// ErrEmptyKey is returned when empty key is provided.
	ErrEmptyKey = errors.New("empty key provided")
	// ErrNotFound is returned when key is not found in map.
	ErrNotFound = errors.New("key not found")
	// ErrNilPartitioner is returned when partitioner is nil.
	ErrNilPartitioner = errors.New("partitioner must not be nil")
	// ErrInvalidPartitions is returned when partitions number is zero.
	ErrInvalidPartitions = errors.New("number of partitions must be greater than zero")
)
