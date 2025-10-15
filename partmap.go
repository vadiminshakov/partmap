package partmap

type partitioner interface {
	// Find returns the partition index for the provided key.
	Find(key string) uint
}

// PartitionedMap is a map that is partitioned into several maps.
type PartitionedMap struct {
	partsnum   uint
	partitions []*partition
	finder     partitioner
}

// NewPartitionedMap creates new partitioned map with given partitioner and number of partitions.
func NewPartitionedMap(partitioner partitioner, partsnum uint, partSize uint) (*PartitionedMap, error) {
	if partitioner == nil {
		return nil, ErrNilPartitioner
	}
	if partsnum == 0 {
		return nil, ErrInvalidPartitions
	}

	partitions := makePartitions(partsnum, partSize)

	return &PartitionedMap{partsnum: partsnum, partitions: partitions, finder: partitioner}, nil
}

// NewPartitionedMapWithDefaultPartitioner creates new partitioned map with default partitioner and given number of partitions.
func NewPartitionedMapWithDefaultPartitioner(partsnum uint, partSize uint) (*PartitionedMap, error) {
	partitioner, err := NewHashSumPartitioner(partsnum)
	if err != nil {
		return nil, err
	}

	return NewPartitionedMap(partitioner, partsnum, partSize)
}

func makePartitions(partsnum uint, partSize uint) []*partition {
	partitions := make([]*partition, 0, partsnum)
	for i := 0; i < int(partsnum); i++ {
		m := make(map[string]any, partSize)
		partitions = append(partitions, &partition{stor: m})
	}
	return partitions
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
