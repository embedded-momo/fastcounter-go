# fastcounter-go

A fast concurrent counter for write-heavy workload.

## Benchmark

For write-heavy workload, `Add` is on par with [chen3feng/atomiccounter](https://github.com/chen3feng/atomiccounter), while the `Read` is faster.

AMD Ryzen 7 5800H:

```
goos: linux
goarch: amd64
pkg: atomiccounter_bench
cpu: AMD Ryzen 7 5800H with Radeon Graphics
                      │   result    │
                      │   sec/op    │
Add_NonAtomic-16        941.8n ± 1%
Add_Atomic-16           1.613µ ± 0%
Add_AtomicCounter-16    118.2n ± 6%
Add_XsyncCounter-16     198.0n ± 4%
Add_GoAdder-16          407.6n ± 1%
Add_GarrAdder-16        404.1n ± 1%
Add_FastCounter-16      115.3n ± 0%

Read_NonAtomic-16       4.879n ± 0%
Read_Atomic-16          4.894n ± 0%
Read_AtomicCounter-16   402.5n ± 1%
Read_XSyncCounter-16    337.5n ± 0%
Read_GoAdder-16         33.51n ± 0%
Read_GarrAdder-16       34.32n ± 0%
Read_FastCounter-16     288.0n ± 0%

geomean                 135.8n
```

Apple M1 Pro:

```
goos: darwin
goarch: arm64
pkg: atomiccounter_bench
cpu: Apple M1 Pro
                      │    result    │
                      │    sec/op    │
Add_NonAtomic-10        21.58n ±  0%
Add_Atomic-10           7.548µ ±  1%
Add_AtomicCounter-10    98.88n ± 16%
Add_XsyncCounter-10     161.0n ±  3%
Add_GoAdder-10          1.950µ ±  7%
Add_GarrAdder-10        1.404µ ±  3%
Add_FastCounter-10      97.55n ±  1%

Read_NonAtomic-10       4.051n ±  0%
Read_Atomic-10          8.407n ±  6%
Read_AtomicCounter-10   277.4n ±  0%
Read_XSyncCounter-10    467.4n ±  0%
Read_GoAdder-10         34.21n ±  1%
Read_GarrAdder-10       34.16n ±  1%
Read_FastCounter-10     130.6n ±  6%

geomean                 131.6n
```

## How does it works?

fastcounter follows the same approach as [chen3feng/atomiccounter](https://github.com/chen3feng/atomiccounter), with minor simplification:

- fastcounter doesn't allocate in chunks, and can't save memory by sharing cells
- fastcounter use a more pessimistic guess of cache line size, at the cost of wasting memory
- `Read` is manually unrolled to read two int64s in one go
- code is simple, you can easily adopt it to your actual need (different cache size, different shards, etc)
 
## Should I use this package?

No, this package is mainly for learning purpose.

Please consider `sync/atomic`first, unless you have high frequency/high concurrency operations and atomic operations become the bottleneck.
