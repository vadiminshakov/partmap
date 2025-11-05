package partmap

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

type XXHashPartitioner struct {
	partitionsNum uint
	mask          uint64
	useMask       bool
}

func NewXXHashPartitioner(partitionsNum uint) (*XXHashPartitioner, error) {
	if partitionsNum == 0 {
		return nil, fmt.Errorf("partmap: %w", ErrInvalidPartitions)
	}

	// power-of-two fast path
	useMask := (partitionsNum & (partitionsNum - 1)) == 0
	var mask uint64
	if useMask {
		mask = uint64(partitionsNum - 1)
	}

	return &XXHashPartitioner{
		partitionsNum: partitionsNum,
		mask:          mask,
		useMask:       useMask,
	}, nil
}

func (h *XXHashPartitioner) Find(key string) uint {
	sum := xxhash.Sum64String(key)
	if h.useMask {
		return uint(sum & h.mask)
	}
	
	return uint(sum % uint64(h.partitionsNum))
}