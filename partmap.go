package partmap

type partitioner interface {
	Find(key string) uint
}

// PartitionedMap is a map that is partitioned into several maps.
type PartitionedMap struct {
	partsnum   uint
	partitions []*partition
	finder     partitioner
}

// NewPartitionedMap creates new partitioned map with given partitioner and number of partitions.
func NewPartitionedMap(partitioner partitioner, partsnum uint, partSize uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any, partSize)
		partitions = append(partitions, &partition{stor: m})
	}
	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: partitioner}
}

// NewPartitionedMapWithDefaultPartitioner creates new partitioned map with default partitioner and given number of partitions.
func NewPartitionedMapWithDefaultPartitioner(partsnum uint, partSize uint) *PartitionedMap {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any, partSize)
		partitions = append(partitions, &partition{stor: m})
	}

	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: NewHashSumPartitioner(partsnum)}
}

// Set sets value for given key.
func (c *PartitionedMap) Set(key string, value any) error {
	if len(key) == 0 {
		return ErrEmptyKey
	}

	partitionIndex := c.finder.Find(key)
	part := c.partitions[partitionIndex]

	part.set(key, value)

	return nil
}

// Get returns value for given key.
func (c *PartitionedMap) Get(key string) (any, error) {
	if len(key) == 0 {
		var emptyT any
		return emptyT, ErrEmptyKey
	}

	partitionIndex := c.finder.Find(key)
	part := c.partitions[partitionIndex]

	value, ok := part.get(key)
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

	partitionIndex := c.finder.Find(key)
	part := c.partitions[partitionIndex]

	part.del(key)

	return nil
}
