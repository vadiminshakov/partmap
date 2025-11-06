package partmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type stubPartitioner struct {
	idx uint
}

func (s stubPartitioner) Find(string) uint {
	return s.idx
}

func TestSetGetDel(t *testing.T) {
	m, err := NewPartitionedMapWithDefaultPartitioner(1000, 10)
	require.NoError(t, err)
	require.NoError(t, m.Set("1", 1))
	v, getErr := m.Get("1")
	require.NoError(t, getErr)
	require.Equal(t, 1, v)

	require.NoError(t, m.Del("1"))
	v, getErr = m.Get("1")
	require.ErrorIs(t, getErr, ErrNotFound)
	require.Nil(t, v)
}

func TestSetEmptyKey(t *testing.T) {
	m, err := NewPartitionedMapWithDefaultPartitioner(1000, 10)
	require.NoError(t, err)
	require.ErrorIs(t, m.Set("", 1), ErrEmptyKey)
}

func TestGetEmptyKey(t *testing.T) {
	m, err := NewPartitionedMapWithDefaultPartitioner(1000, 10)
	require.NoError(t, err)
	_, getErr := m.Get("")
	require.ErrorIs(t, getErr, ErrEmptyKey)
}

func TestDelEmptyKey(t *testing.T) {
	m, err := NewPartitionedMapWithDefaultPartitioner(1000, 10)
	require.NoError(t, err)
	require.ErrorIs(t, m.Del(""), ErrEmptyKey)
}

func TestNewPartitionedMap_ReturnsErrorOnZeroPartitions(t *testing.T) {
	m, err := NewPartitionedMap(stubPartitioner{}, 0, 10)
	require.Nil(t, m)
	require.ErrorIs(t, err, ErrInvalidPartitions)
}

func TestNewPartitionedMap_ReturnsErrorOnNilPartitioner(t *testing.T) {
	m, err := NewPartitionedMap(nil, 1, 10)
	require.Nil(t, m)
	require.ErrorIs(t, err, ErrNilPartitioner)
}

func TestXXHashPartitioner_ReturnsErrorOnZeroPartitions(t *testing.T) {
	p, err := NewXXHashPartitioner(0)
	require.Nil(t, p)
	require.ErrorIs(t, err, ErrInvalidPartitions)
}
