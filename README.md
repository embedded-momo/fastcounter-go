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
Add_NonAtomic-16        924.3n ± 0%
Add_Atomic-16           1.608µ ± 0%
Add_AtomicCounter-16    117.9n ± 8%
Add_XsyncCounter-16     194.1n ± 4%
Add_GoAdder-16          408.1n ± 1%
Add_GarrAdder-16        405.8n ± 1%
Add_FastCounter-16      115.5n ± 9%

Read_NonAtomic-16       4.877n ± 0%
Read_Atomic-16          4.893n ± 0%
Read_AtomicCounter-16   401.8n ± 1%
Read_XSyncCounter-16    337.2n ± 0%
Read_GoAdder-16         33.47n ± 0%
Read_GarrAdder-16       34.31n ± 0%
Read_FastCounter-16     148.4n ± 0%

geomean                 129.1n

```

Apple M1 Pro:

```
goos: darwin
goarch: arm64
pkg: atomiccounter_bench
cpu: Apple M1 Pro
                      │    result    │
                      │    sec/op    │
Add_NonAtomic-10        21.82n ±  0%
Add_Atomic-10           7.398µ ±  2%
Add_AtomicCounter-10    99.44n ± 16%
Add_XsyncCounter-10     153.7n ± 10%
Add_GoAdder-10          1.951µ ±  7%
Add_GarrAdder-10        1.412µ ±  4%
Add_FastCounter-10      97.80n ±  0%

Read_NonAtomic-10       4.045n ±  0%
Read_Atomic-10          7.940n ±  1%
Read_AtomicCounter-10   264.1n ±  0%
Read_XSyncCounter-10    441.9n ±  0%
Read_GoAdder-10         33.69n ±  1%
Read_GarrAdder-10       33.55n ±  1%
Read_FastCounter-10     73.17n ±  0%

geomean                 124.1n
```

## How does it works?

fastcounter follows the same approach as [chen3feng/atomiccounter](https://github.com/chen3feng/atomiccounter), with minor simplification:

- fastcounter doesn't allocate in chunks, and can't save memory by sharing cells
- fastcounter use a more pessimistic guess of cache line size, at the cost of wasting memory
- `Read` is manually unrolled to read four int64s in one go
- code is simple, you can easily adopt it to your actual need (different cache size, different shards, etc)
 
## Should I use this package?

No, this package is mainly for learning purpose.

Please consider `sync/atomic`first, unless you have high frequency/high concurrency operations and atomic operations become the bottleneck.
