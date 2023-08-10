package main

import (
	"errors"
	"hash/crc32"
	"sync"
)

type Partitioner interface {
	Find(key string) (uint, error)
}

type hashSumPartitioner struct {
	partitionsNum uint
}

func (h *hashSumPartitioner) Find(key string) (uint, error) {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitionsNum, nil
}

type Cache struct {
	n          uint
	partitions []*partition
	finder     Partitioner
}

type partition struct {
	stor map[string]any
	sync.RWMutex
}

func (p *partition) set(key string, value any) {
	p.Lock()
	p.stor[key] = value
	p.Unlock()
}

func (p *partition) get(key string) (any, bool) {
	p.RLock()
	v, ok := p.stor[key]
	if !ok {
		p.RUnlock()
		return nil, false
	}
	p.RUnlock()
	return v, true
}

func NewCache(partitioner Partitioner, shardsNum uint) *Cache {
	partitions := make([]*partition, 0, shardsNum)
	for i := 0; i < int(shardsNum); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{stor: m})
	}
	return &Cache{n: shardsNum, partitions: partitions, finder: partitioner}
}

func (c *Cache) Set(key string, value any) error {
	partitionIndex, err := c.finder.Find(key)
	if err != nil {
		return err
	}

	partition := c.partitions[partitionIndex]
	partition.set(key, value)

	return nil
}
func (c *Cache) Get(key string) (any, error) {
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

func main() {
	cache := NewCache(&hashSumPartitioner{partitionsNum: 100}, 100)

	cache.Set("1", 1)

	cache.Set("2", 2)
}
