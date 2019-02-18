package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFun(t *testing.T) {
	s := "A,B,C,D"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvFun(t *testing.T) {
	var s string = strconv.Itoa(1)
	var (
		i   int
		err error
	)
	if i, err = strconv.Atoi(s); err == nil {
		t.Logf("%d", i)
	} else {
		t.Error(err)
	}
}
