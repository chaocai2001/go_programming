package concurrency

import (
	"testing"
	"time"
)

func SlowService() string {
	//time.Sleep(10 * time.Second)
	return "Done."
}

func TestTimeout(t *testing.T) {
	retCh := make(chan string)
	go func() {
		retCh <- SlowService()
	}()
	select {
	case ret := <-retCh:
		t.Logf("result %s", ret)
	case <-time.After(time.Second * 1):
		t.Error("time out")
	}
}
