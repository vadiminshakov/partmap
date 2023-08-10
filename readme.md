**Simple and fast partitioned map**

Faster then writing to std map
```
go test -bench=. -benchtime=3s

goos: darwin
goarch: amd64
pkg: github.com/vadimInshakov/partmap
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkStd/set_std_concurrently-12                  3289076   1332 ns/op
BenchmarkSyncStd/set_sync_map_std_concurrently-12     2408612   1691 ns/op
BenchmarkPartitioned/set_partitioned_concurrently-12  13536134  408.6 ns/op <-
```

Usage:

```
m := partmap.NewPartitionedMapWithDefaultPartitioner(3)
m.Set("key", 123)
value, _ := m.Get("key")
```

