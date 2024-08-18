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
	c1 := int64(0)
	c2 := int64(0)
	c3 := int64(0)
	c4 := int64(0)

	for i := 0; i < 64; i += 4 {
		c1 += c.counters[i].value
		c2 += c.counters[i+1].value
		c3 += c.counters[i+2].value
		c4 += c.counters[i+3].value
	}

	return c1 + c2 + c3 + c4
}
