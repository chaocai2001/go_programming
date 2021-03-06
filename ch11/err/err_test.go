package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError error = errors.New("n must be greater than 2")
var GreaterThanHundredError error = errors.New("n must be less than 100")

func GetFibonacci(n int) ([]int, error) {
	fibList := []int{1, 1}
	// if n < 2 || n > 100 {
	// 	return nil, errors.New("n must be in the range [0,100]")
	// }
	if n < 0 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, GreaterThanHundredError
	}

	for i := 2; /*短变量声明 := */ i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {

		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}

func TestGetFibonacci(t *testing.T) {
	var list []int
	list, err := GetFibonacci(10)
	if err == LessThanTwoError {
		t.Error("Need a larger number")
	}

	if err == GreaterThanHundredError {
		t.Error("Need a larger number")
	}
	if err != nil {
		t.Error(err)
	}

	t.Log(list)

	GetFibonacci1("5")
	GetFibonacci2("5")
}
