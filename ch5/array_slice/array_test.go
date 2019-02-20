package array_slice

import (
	"fmt"
	"testing"
)

func TestArrayDeclaration(t *testing.T) {
	var a [3]int //声明并初始化为默认零值
	a[0] = 1

	b := [3]int{1, 2, 3}           //声明同时初始化
	c := [2][2]int{{1, 2}, {3, 4}} //多维数组初始化
	t.Log(a, b, c)
}

func TestTravelArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	for idx, elem := range a {
		fmt.Println(idx, elem)
	}
	fmt.Println(a[1:])
	fmt.Println(a[:3])
}
