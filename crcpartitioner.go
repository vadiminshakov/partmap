package partmap

import (
	"fmt"
	"hash/crc32"
)

type HashSumPartitioner struct {
	partitionsNum uint
}

func NewHashSumPartitioner(partitionsNum uint) (*HashSumPartitioner, error) {
	if partitionsNum == 0 {
		return nil, fmt.Errorf("partmap: %w", ErrInvalidPartitions)
	}

	return &HashSumPartitioner{partitionsNum: partitionsNum}, nil
}

func (h *HashSumPartitioner) Find(key string) uint {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitionsNum
}
