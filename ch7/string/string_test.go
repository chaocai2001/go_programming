package string_test

import (
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值“”
	s = "hello"
	t.Log(len(s))
	//	s[1] = '3'
	s = "\xE4\xB8\xA5"
	s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s))
	t.Logf("%x", s)
	c := []rune(s)
	t.Logf("%x", c[0])
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		t.Logf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
