package basic_type

import (
	"math"
	"testing"
)

type MyByte uint8

func TestTypeImplicitConv(t *testing.T) {
	var (
		i   int
		err error
	)
	//var i int32 = 16
	//var j int64 = i
	//t.Log(i, j)
}

func TestTypeAlias(t *testing.T) {
	var i MyByte = 8
	j := uint8(i)
	t.Log(j)
	//var k uint8 = i
}

func TestMathDef(t *testing.T) {
	var f float64
	t.Log(math.MaxUint32, f)
	t.Log(math.IsNaN(f / 0))
}

func TestPtr(t *testing.T) {
	i := 10
	var j *int = &i
	*j = 1000
	t.Log(i)
	//	j++
}

func TestInitString(t *testing.T) {
	var s string
	t.Log(s == "")
	i := 0
	i++
}

func TestCompareArrays(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [...]int{1, 2, 4}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(a == c)
	//t.Log(a != d)
}
