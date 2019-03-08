package testing

import "testing"

func Foo() *int {
	var i int
	i = 10
	return &i
}

func TestFoo(t *testing.T) {
	x := Foo()
	t.Log(*x)
}
