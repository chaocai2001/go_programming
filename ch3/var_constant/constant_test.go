package var_constant

import (
	"fmt"
	"testing"
)

const (
	Monday = 2
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestContinuousConstants(t *testing.T) {
	fmt.Println(Monday, Sunday)
	fmt.Printf("%b,%b,%b", Readable, Writable, Executable)
}

func TestBitOp(t *testing.T) {
	a := 7
	t.Logf("%b %t %t %t", a, a&Readable == Readable, a&Writable == Writable,
		a&Executable == Executable)
	//clean Readable
	a = a &^ Readable
	t.Logf("%b %t %t %t", a, a&Readable == Readable, a&Writable == Writable,
		a&Executable == Executable)
}
