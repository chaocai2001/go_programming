package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m, m["one"])
	m1 := map[string]int{}
	m1["one"] = 1
	t.Log(m1)
	m2 := make(map[string]int, 10)
	t.Log(len(m2))
}

func TestAccessEmptyKey(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log("four", m["four"])
	if v, ok := m["four"]; ok {
		t.Log("four", v)
	} else {
		t.Log("Not existing")
	}

}

func TestMapTravel(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}
