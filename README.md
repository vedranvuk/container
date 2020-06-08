# container

WARNING: Work in progress. APIs might change but tags will be added to discern versions if it does. If using in production, please use vendoring or repackage into your own package.

## Description


Various containers and datastructures.

* queue
  * fifo: First-In-First-Out queue.
  * ttl: Timeout list of items with Time To Live.


## fifo

A singly linked concurency safe first in first out list with Push, Pop, Peek and Remove operations.

## ttl

A doubly linked concurency safe list with item specific timeouts using an additional map for fast item access. Insert positions sorted by item timeout timestamps are done starting from end of list by default from presumption list will serve for session tracking with all sessions having same timeouts.

## Benchmarks

### fifo

An i5 skylake lap with 8GB RAM.

```
pkg: github.com/vedranvuk/container/queue/fifo
BenchmarkQueuePush
BenchmarkQueuePush-4             7219392               185 ns/op
BenchmarkQueuePop
BenchmarkQueuePop-4             55198738                20.5 ns/op
BenchmarkQueuePeek
BenchmarkQueuePeek-4            64719546                18.5 ns/op
BenchmarkQueueRemoveFirst
BenchmarkQueueRemoveFirst-4      3925570               310 ns/op
BenchmarkQueueRemoveLast
BenchmarkQueueRemoveLast-4         49438            117964 ns/op
BenchmarkQueueRemoveRandom
BenchmarkQueueRemoveRandom-4       84730            130831 ns/op
```

### ttl

Average times with 0 .. 1s random timeouts.

```
pkg: github.com/vedranvuk/container/queue/ttl
BenchmarkTTLAdd
BenchmarkTTLAdd-4                2275051               523 ns/op
BenchmarkTTLRemove
BenchmarkTTLRemove-4             8822643               119 ns/op
BenchmarkTTLRemoveReverse
BenchmarkTTLRemoveReverse-4     34743085                29.8 ns/op
BenchmarkTTLReset
BenchmarkTTLReset-4              6620694               155 ns/op
BenchmarkTTLResetReverse
BenchmarkTTLResetReverse-4       6337597               182 ns/op
```

## License

MIT. See incliuded LICENSE file.