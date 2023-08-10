package partmap

import "hash/crc32"

type HashSumPartitioner struct {
	partitionsNum uint
}

func NewHashSumPartitioner(partitionsNum uint) *HashSumPartitioner {
	return &HashSumPartitioner{partitionsNum: partitionsNum}
}

func (h *HashSumPartitioner) Find(key string) (uint, error) {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitionsNum, nil
}
