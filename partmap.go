package partmap

import (
	"errors"
)

type partitioner interface {
	Find(key string) (uint, error)
}

type PartitionedMap struct {
	partsnum   uint
	partitions []*partition
	finder     partitioner
}

func NewPartitionedMap(partitioner partitioner, partsnum uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{stor: m})
	}
	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: partitioner}
}

func NewPartitionedMapWithDefaultPartitioner(partsnum uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{stor: m})
	}

	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: NewHashSumPartitioner(partsnum)}
}

func (c *PartitionedMap) Set(key string, value any) error {
	if len(key) == 0 {
		return ErrEmptyKey
	}
	var emptyT any
	if value == emptyT {
		return ErrEmptyValue
	}

	partitionIndex, err := c.finder.Find(key)
	if err != nil {
		return err
	}

	partition := c.partitions[partitionIndex]
	partition.set(key, value)

	return nil
}
func (c *PartitionedMap) Get(key string) (any, error) {
	if len(key) == 0 {
		var emptyT any
		return emptyT, ErrEmptyKey
	}

	partitionIndex, err := c.finder.Find(key)
	if err != nil {
		return nil, err
	}

	partition := c.partitions[partitionIndex]
	value, ok := partition.get(key)
	if !ok {
		return nil, errors.New("no such data")
	}

	return value, nil
}
