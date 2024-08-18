package fastcounter

import (
	"sync/atomic"
)

type Counter struct {
	_        [8]int64
	counters [64]counter
}

type counter struct {
	_     [7]int64
	value int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(i int64) {
	idx := threadHash() % 64
	atomic.AddInt64(&c.counters[idx].value, i)
}

func (c *Counter) Read() int64 {
	a := int64(0)
	b := int64(0)

	for i := 0; i < 64; i+=2 {
		a += atomic.LoadInt64(&c.counters[i].value)
		b += atomic.LoadInt64(&c.counters[i+1].value)
	}

	return a + b
}
