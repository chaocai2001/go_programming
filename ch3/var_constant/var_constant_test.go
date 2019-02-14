package var_constant

import "testing"

func TestVarExVal(t *testing.T) {
	a := 1
	b := 2
	tmp := a
	a = b
	b = tmp
	//a, b = b, a
	t.Log(a, b)
}
