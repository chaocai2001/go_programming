package container_test

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	e1 := l.PushFront("Hello")
	l.InsertAfter("World", e1)

	for i := l.Front(); i != nil; i = i.Next() {
		t.Logf("%T %v", i.Value, i.Value)
	}

	for i := l.Back(); i != nil; i = i.Prev() {
		t.Logf("%T %v", i.Value, i.Value)
	}
}
