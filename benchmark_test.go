package fastcounter_test

import (
	"sync/atomic"
	"testing"

	"github.com/embedded-momo/fastcounter-go"
)

func BenchmarkAtomicAdd(b *testing.B) {
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt64(&count, 1)
		}
	})
}

func BenchmarkCounter(b *testing.B) {
	counter := fastcounter.NewCounter()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Add(1)
		}
	})
}

func BenchmarkAtomicRead(b *testing.B) {
	count := int64(0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 100; i++ {
				_ = atomic.LoadInt64(&count)
			}
		}
	})
}

func BenchmarkCounterRead(b *testing.B) {
	counter := fastcounter.NewCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 100; i++ {
				_ = counter.Read()
			}
		}
	})
}
