package partmap

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkStd(b *testing.B) {
	m := make(map[string]int)
	b.Run("set std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m[fmt.Sprint(i)] = i
		}
	})
	b.Run("set std concurrently", func(b *testing.B) {
		var wg sync.WaitGroup
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			i := i
			go func() {
				mu.Lock()
				m[fmt.Sprint(i)] = i
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	})
	b.Run("get std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m[fmt.Sprint(i)]
		}
	})
	b.Run("get std concurrently", func(b *testing.B) {
		var wg sync.WaitGroup
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			i := i
			go func() {
				mu.RLock()
				_, _ = m[fmt.Sprint(i)]
				mu.RUnlock()
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func BenchmarkSyncStd(b *testing.B) {
	b.Run("set sync map std concurrently", func(b *testing.B) {
		var m sync.Map
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			i := i
			go func() {
				m.Store(fmt.Sprint(i), i)
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func BenchmarkPartitioned(b *testing.B) {
	m := NewPartitionedMap(NewHashSumPartitioner(1000), 1000)
	b.Run("set partitioned", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprint(i), i)
		}
	})
	b.Run("set partitioned concurrently", func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			i := i
			go func() {
				m.Set(fmt.Sprint(i), i)
				wg.Done()
			}()
		}
		wg.Wait()
	})
	b.Run("get partitioned", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = m.Get(fmt.Sprint(i))
		}
	})
	b.Run("get partitioned concurrently", func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			i := i
			go func() {
				_, _ = m.Get(fmt.Sprint(i))
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
