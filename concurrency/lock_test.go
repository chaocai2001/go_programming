package concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer mu.Unlock()
			mu.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafeRW(t *testing.T) {
	counter := 0
	var mu sync.RWMutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer mu.Unlock()
			mu.Lock()
			counter++
		}()

		go func() {
			defer mu.RUnlock()
			mu.RLock()
			//t.Log(counter)
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
