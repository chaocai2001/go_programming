package empty_interface_assertion

import (
	"fmt"
	"testing"
)

type Coder interface {
	WriteCode() string
}

type GoCoder struct {
}

func (p *GoCoder) WriteCode() string {
	return "fmt.Println(\"Hello World\")"
}

func doSomething(any interface{}) {
	if i, ok := any.(int); ok {
		fmt.Println("Integer:", i)
		return
	}
	switch v := any.(type) {
	case int:
		fmt.Println("Integer:", v)
	case Coder:
		fmt.Println("Programmer", v.WriteCode())

	}
	fmt.Println(any)
}
func TestEmptyInterface(t *testing.T) {
	i := 1000
	doSomething(i)
	doSomething(&i)
	g := new(GoCoder)
	doSomething(g)
}
