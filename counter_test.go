package fastcounter_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/embedded-momo/fastcounter-go"
)

func TestCounter(t *testing.T) {
	counter := fastcounter.NewCounter()

	m := rand.Int63n(10000)
	n := rand.Int63n(100)

	for i := int64(0); i < m; i++ {
		counter.Add(n)
	}

	if counter.Read() != m*n {
		t.Errorf("expect count %d, got: %d", m*n, counter.Read())
	}
}

func TestCounterConcurrent(t *testing.T) {
	counter := fastcounter.NewCounter()

	m := rand.Int63n(10000)
	n := rand.Int63n(100)

	wg := sync.WaitGroup{}

	for k := 0; k < 10; k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := int64(0); i < m; i++ {
				counter.Add(n)
			}
		}()
	}

	wg.Wait()

	if counter.Read() != m*n*10 {
		t.Errorf("expect count %d, got: %d", m*n, counter.Read())
	}
}
