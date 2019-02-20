package demos

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	n := 0
	for n < 10 {
		n++
		fmt.Println(n)
	}
}
