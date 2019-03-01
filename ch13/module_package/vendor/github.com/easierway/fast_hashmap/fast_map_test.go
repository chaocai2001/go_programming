package fast_hashmap

import "testing"

func mapBasicTest(m MapWithStringKey, t *testing.T) {

	m.Put("100", 100)
	v, ok := m.Get("100")
	if !ok {
		t.Error("Failed to get")
	}
	vInt, ok := v.(int)
	if !ok {
		t.Error("Failed to get")
	}
	if vInt != 100 {
		t.Errorf("Wrong value %v", v)
	}
	v, ok = m.Get("200")
	if ok {
		t.Error("Failed to get when no such key")
	}
	m.Put("100", 200)
	v, ok = m.Get("100")
	if !ok {
		t.Error("Failed to get")
	}
	vInt, ok = v.(int)
	if !ok {
		t.Error("Failed to get")
	}
	if vInt != 200 {
		t.Errorf("Wrong value %v", v)
	}
}

func TestFastMap(t *testing.T) {
	m := NewFastMap(100)
	mapBasicTest(m, t)
}

func TestFasterMap(t *testing.T) {
	m := NewFasterMap(100)
	mapBasicTest(m, t)
}
