package array_slice

import "testing"

func TestInitSlice(t *testing.T) {
	var s0 []int
	t.Log("s0", len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log("s0", len(s0), cap(s0))
	s := []int{}
	t.Log(len(s), cap(s))

	s1 := []int{1, 2, 3}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 2, 4)
	t.Log(s2[0], s2[1], len(s2), cap(s2))
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	Summer := year[5:8]
	t.Log(Summer, len(Summer), cap(Summer))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, 1)
		t.Log(len(s), cap(s))
	}

}

// func TestMap(t *testing.T) {
// 	m := make(map[int]interface{}, 10)
// 	m[1] = nil
// 	if v, ok := m[1]; ok {
//     if
// 		t.Log(v)
// 	}
// 	t.Log("Empty")
// }
