package map_test

import (
	"testing"
)

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

func TestFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2))
}

func TestUsingMapAsSet(t *testing.T) {
	set := map[int]bool{}
	//添加元素
	set[1] = true
	//删除元素
	delete(set, 1)
	t.Log(len(set))
	for i := 0; i < 3; i++ {
		//判断元素是否存在
		if set[i] {
			t.Logf("%d is existing.", i)
		}
	}

}
