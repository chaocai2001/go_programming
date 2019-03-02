package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i, input := range inputs {
		ret := square(input)
		if ret != expected[i] {
			t.Errorf("The expected value is %d, actual value is %d", expected[i], ret)
		}
	}
}

func TestFatal(t *testing.T) {
	fmt.Println("start")
	t.Fatal("err")
	fmt.Println("end")
}

func TestError(t *testing.T) {
	fmt.Println("start")
	t.Error("err")
	fmt.Println("end")
}

func TestCube(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 8, 27}
	for i, input := range inputs {
		ret := cube(input)
		if ret != expected[i] {
			t.Errorf("The expected value is %d, actual value is %d", expected[i], ret)
		}
	}
}

func TestSquareWithAssertion(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	assert := assert.New(t)
	for i, input := range inputs {
		ret := square(input)
		assert.Equal(expected[i], ret)

	}
}
