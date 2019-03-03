package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name       string
	Age        int
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 4: "three"}
	//fmt.Println(a == b)
	fmt.Println(reflect.DeepEqual(a, b))
	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

func CheckType(v interface{}) {
	t := reflect.ValueOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)

}

func fillNameAndAge(st interface{}, settings map[string]interface{}) {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	for k, v := range settings {
		if field, ok := (reflect.ValueOf(st)).Elem().Type().FieldByName(k); ok {
			if field.Type == reflect.TypeOf(v) {
				fmt.Println(field.Type, reflect.TypeOf(v))
				vstr := reflect.ValueOf(st)
				vstr = vstr.Elem()
				vstr.FieldByName(k).Set(reflect.ValueOf(v))
			}
		}

	}
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	fillNameAndAge(&e, settings)
	t.Log(e)
	c := new(Customer)
	fillNameAndAge(c, settings)
	t.Log(*c)

}
