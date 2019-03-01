package container

import (
	"container/heap"
	"container/list"
	"testing"
)

func TestHeap(t *testing.T) {
	l := list.New()
	h := heap.Init(l)
}
