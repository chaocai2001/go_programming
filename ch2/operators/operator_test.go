package operators

import "testing"

func TestCompareArrays(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [...]int{1, 2, 4}
	//	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
	//t.Log(a != d)
}

func TestClearByBit(t *testing.T) {
	a := 0xFF00
	b := 0xFFFF
	c := 0x0FFF
	t.Log(a &^ b)
	t.Logf("%x", a&^c)
}
