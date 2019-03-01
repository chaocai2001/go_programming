package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestGorouting(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
