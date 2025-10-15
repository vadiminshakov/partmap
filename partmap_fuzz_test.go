package partmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzPartMap(f *testing.F) {
	m, err := NewPartitionedMapWithDefaultPartitioner(1000, 10)
	if err != nil {
		f.Fatalf("failed to create partitioned map: %v", err)
	}

	f.Add("key", 1)
	f.Add("1", 2)
	f.Fuzz(func(t *testing.T, key string, value int) {
		err := m.Set(key, value)
		if err == ErrEmptyKey {
			t.Skip()
		}
		require.NoError(t, err)

		v, err := m.Get(key)
		if err == ErrEmptyKey {
			t.Skip()
		}
		require.NoError(t, err)
		require.Equal(t, value, v)
	})
}
