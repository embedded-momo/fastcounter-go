package fastcounter

import (
	"sync/atomic"
)

type Counter struct {
	counters [64]counter
}

type counter struct {
	value int64
	_     [64 - 8]byte
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(i int64) {
	if c == nil {
		panic("nil Counter")
	}
	idx := threadHash() % 64
	atomic.AddInt64(&c.counters[idx].value, i)
}

func (c *Counter) Read() int64 {
	if c == nil {
		panic("nil Counter")
	}

	a := int64(0)
	b := int64(0)
	e := int64(0)
	d := int64(0)

	for i := 0; i < 64; i += 8 {
		a += atomic.LoadInt64(&c.counters[i+0].value)
		b += atomic.LoadInt64(&c.counters[i+1].value)
		e += atomic.LoadInt64(&c.counters[i+2].value)
		d += atomic.LoadInt64(&c.counters[i+3].value)

		a += atomic.LoadInt64(&c.counters[i+4].value)
		b += atomic.LoadInt64(&c.counters[i+5].value)
		e += atomic.LoadInt64(&c.counters[i+6].value)
		d += atomic.LoadInt64(&c.counters[i+7].value)
	}
	return a + b + e + d
}
