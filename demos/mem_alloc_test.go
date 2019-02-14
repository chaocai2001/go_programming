package demos

import (
	"fmt"
	"testing"
)

func foo() *string {
	s := "Hello World"
	return &s
}

func TestFoo(t *testing.T) {
	fmt.Println(*foo())
}
