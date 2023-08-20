![](https://github.com/vadiminshakov/partmap/workflows/tests/badge.svg)

**Simple and fast partitioned map**

Faster then writing to std map.

[Project motivation](https://medium.com/stackademic/writing-a-partitioned-cache-using-go-map-x3-faster-than-the-standard-map-dbfe704fe4bf)

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

**Usage:**

```
m := partmap.NewPartitionedMapWithDefaultPartitioner(3) // 3 partitions
m.Set("key", 123)
value, err := m.Get("key")
if err != nil && !errors.Is(err, partmap.ErrNotFound) {
    panic(err)
}

println(value) // 123

if err := m.Del("key"); err != nil {
    panic(err)
}
```

[https://pkg.go.dev/github.com/vadiminshakov/partmap](https://pkg.go.dev/github.com/vadiminshakov/partmap)