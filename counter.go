package fastcounter

import (
	"sync/atomic"
)

// cacheLineSize is the most pessimistic guess of actual cache line size, at the cost of wasting more memory.
const cacheLineSize = 128

type Counter struct {
	_        [cacheLineSize]int8 // why this is needed?
	counters [64]counter
}

type counter struct {
	value int64
	_     [cacheLineSize - 8]int8
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

	for i := 0; i < 64; i += 2 {
		a += c.counters[i].value
		b += c.counters[i+1].value
	}

	return a + b
}
