package series

import "fmt"

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init1")
}

func square(op int) int {
	return op * op
}

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
