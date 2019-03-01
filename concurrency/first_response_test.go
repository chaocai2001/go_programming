package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func FirstResponse() string {
	numOfRunner := 10
	retCh := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(id int) {
			time.Sleep(10 * time.Millisecond)
			retCh <- fmt.Sprintf("result from %d", id)
		}(i)
	}
	return <-retCh
}

func TestFirstResponse(t *testing.T) {
	t.Log(FirstResponse())
}
