package partmap

type partitioner interface {
	Find(key string) (uint, error)
}

// PartitionedMap is a map that is partitioned into several maps.
type PartitionedMap struct {
	partsnum   uint
	partitions []*partition
	finder     partitioner
}

// NewPartitionedMap creates new partitioned map with given partitioner and number of partitions.
func NewPartitionedMap(partitioner partitioner, partsnum uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{stor: m})
	}
	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: partitioner}
}

// NewPartitionedMapWithDefaultPartitioner creates new partitioned map with default partitioner and given number of partitions.
func NewPartitionedMapWithDefaultPartitioner(partsnum uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{stor: m})
	}

	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: NewHashSumPartitioner(partsnum)}
}

// Set sets value for given key.
func (c *PartitionedMap) Set(key string, value any) error {
	if len(key) == 0 {
		return ErrEmptyKey
	}

	partitionIndex, err := c.finder.Find(key)
	if err != nil {
		return err
	}

	partition := c.partitions[partitionIndex]
	partition.set(key, value)

	return nil
}

// Get returns value for given key.
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
		return nil, ErrNotFound
	}

	return value, nil
}

// Del deletes value for given key.
func (c *PartitionedMap) Del(key string) error {
	if len(key) == 0 {
		return ErrEmptyKey
	}

	partitionIndex, err := c.finder.Find(key)
	if err != nil {
		return err
	}

	partition := c.partitions[partitionIndex]
	partition.del(key)

	return nil
}
