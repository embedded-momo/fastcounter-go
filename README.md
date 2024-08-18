# fastcounter-go

A fast concurrent counter for X86-64.

## Benchmark

For write-heavy workload, it's on par with [chen3feng/atomiccounter](https://github.com/chen3feng/atomiccounter), while the read path is faster then previous one.

With 16 threads:

```
goos: linux
goarch: amd64
pkg: atomiccounter_bench
cpu: AMD Ryzen 7 5800H with Radeon Graphics
                      │   result    │
                      │   sec/op    │
Add_NonAtomic-16        946.3n ± 0%
Add_Atomic-16           1.612µ ± 0%
Add_AtomicCounter-16    118.8n ± 1%
Add_XsyncCounter-16     253.5n ± 1%
Add_GoAdder-16          409.6n ± 1%
Add_GarrAdder-16        405.0n ± 0%
Add_FastCounter-16      116.5n ± 8%

Read_NonAtomic-16       4.884n ± 0%
Read_Atomic-16          4.881n ± 0%
Read_AtomicCounter-16   401.8n ± 1%
Read_XSyncCounter-16    337.8n ± 0%
Read_GoAdder-16         33.50n ± 0%
Read_GarrAdder-16       34.34n ± 0%
Read_FastCounter-16     348.2n ± 0%
```

With 2 threads:

```
goos: linux
goarch: amd64
pkg: atomiccounter_bench
cpu: AMD Ryzen 7 5800H with Radeon Graphics
                     │   result    │
                     │   sec/op    │
Add_NonAtomic-2        347.7n ± 1%
Add_Atomic-2           1.037µ ± 0%
Add_AtomicCounter-2    865.1n ± 0%
Add_XsyncCounter-2     1.346µ ± 0%
Add_GoAdder-2          2.010µ ± 0%
Add_GarrAdder-2        2.006µ ± 1%
Add_FastCounter-2      845.1n ± 0%

Read_NonAtomic-2       38.96n ± 0%
Read_Atomic-2          38.88n ± 0%
Read_AtomicCounter-2   3.171µ ± 1%
Read_XSyncCounter-2    2.696µ ± 0%
Read_GoAdder-2         267.3n ± 0%
Read_GarrAdder-2       274.2n ± 0%
Read_FastCounter-2     2.780µ ± 0%
```

## How does it works?

fastcounter follows the same approach as [chen3feng/atomiccounter](https://github.com/chen3feng/atomiccounter), with minor simplification:

- it doesn't allocate in chunks, and can't save memory by sharing cells
- the `Read` loop is manually unrolled to read two int64s in one go
- `Read` utlizes atomic read, this makas no difference for X86-64, but it's much slower for ARM64.
 
## Should I use this package?

No, this package is mainly for learning purpose.

Please consider `sync/atomic`first, unless you have high frequency and high concurrency operations.
