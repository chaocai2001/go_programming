package array_slice

import "testing"

func TestComparable(t *testing.T) {
	a1 := [...]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 2}
	var s3 []int
	t.Log(a1 == a2)
	t.Log(s1 == nil)
	s1 = append(s1, 10)
	t.Log(cap(s1))

	t.Log(s2 == nil)
	t.Log(s3 == nil)
}
