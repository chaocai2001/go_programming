package fibonacci

import (
	"fmt"
	"testing"
)

const N int = 5 //常量声明

func TestFibonacciCreation(t *testing.T) {
	// 变量声明
	var (
		a int = 1
		b     = 1 //自动类型推断
	)
	// var a int = 1
	// var b = 1
	// var a, b = 1, 1
	fmt.Print(a)
	for i := 0; /*短变量声明 := */ i < N; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + b
	}
	fmt.Println()
}
