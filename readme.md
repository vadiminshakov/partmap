Simple and fast partitioned map.

Usage:

```
m := partmap.NewPartitionedMapWithDefaultPartitioner(3)
m.Set("key", 123)
value, _ := m.Get("1")
```