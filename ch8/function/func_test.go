package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiNum() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func TestFunc(t *testing.T) {
	a, b := returnMultiNum()
	t.Log(a, b)
	timeSpent(func(n int) int {
		time.Sleep(time.Duration(n) * time.Second)
		return 1
	})(1)
	t.Log("sum", sum(1, 2, 3, 4, 5))
}

func timeSpent(inner func(n int) int) func(n int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time escaped ", time.Since(start).Seconds())
		return ret
	}
}

func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Started")
	panic("Fatal error")
}
