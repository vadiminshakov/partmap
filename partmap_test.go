package partmap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetGetDel(t *testing.T) {
	m := NewPartitionedMap(NewHashSumPartitioner(1000), 1000)
	require.NoError(t, m.Set("1", 1))
	v, err := m.Get("1")
	require.NoError(t, err)
	require.Equal(t, 1, v)

	require.NoError(t, m.Del("1"))
	v, err = m.Get("1")
	require.ErrorIs(t, err, ErrNotFound)
	require.Nil(t, v)
}

func TestSetEmptyKey(t *testing.T) {
	m := NewPartitionedMap(NewHashSumPartitioner(1000), 1000)
	require.ErrorIs(t, m.Set("", 1), ErrEmptyKey)
}

func TestGetEmptyKey(t *testing.T) {
	m := NewPartitionedMap(NewHashSumPartitioner(1000), 1000)
	_, err := m.Get("")
	require.ErrorIs(t, err, ErrEmptyKey)
}

func TestDelEmptyKey(t *testing.T) {
	m := NewPartitionedMap(NewHashSumPartitioner(1000), 1000)
	require.ErrorIs(t, m.Del(""), ErrEmptyKey)
}
