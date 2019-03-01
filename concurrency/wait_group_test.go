package concurrency

import (
	"sync"
	"testing"
)

func TestCounterThreadSafeWG(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()
			mu.Lock()
			counter++
		}()
	}
	// time.Sleep(1 * time.Second)
	wg.Wait()
	t.Logf("counter = %d", counter)
}
