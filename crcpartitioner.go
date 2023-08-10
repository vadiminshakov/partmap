package main

import "hash/crc32"

type hashSumPartitioner struct {
	partitionsNum uint
}

func (h *hashSumPartitioner) Find(key string) (uint, error) {
	hashSum := crc32.ChecksumIEEE([]byte(key))

	return uint(hashSum) % h.partitionsNum, nil
}
