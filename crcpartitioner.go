package partmap

import "hash/crc32"

type HashSumPartitioner struct {
	partitionsNum uint
}

func NewHashSumPartitioner(partitionsNum uint) *HashSumPartitioner {
	if partitionsNum == 0 {
		panic("partmap: partitions number must be greater than zero")
	}
	
	return &HashSumPartitioner{partitionsNum: partitionsNum}
}

func (h *HashSumPartitioner) Find(key string) uint {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitionsNum
}
