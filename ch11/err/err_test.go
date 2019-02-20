package err_test

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	if n < 0 || n > 100 {
		return nil, errors.New("n must be in the range [0,100]")
	}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {

}
