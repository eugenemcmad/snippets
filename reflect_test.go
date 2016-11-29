package tests

import (
	"fmt"
	"reflect"
	"testing"
)

func CloneValue(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		destin = x.Interface()
	}
}

func TestTypeSwitch(t *testing.T) {
	var o interface{}
	ss := []string{"s", "s"}

	o = ss
	switch x := o.(type) {
	default:
		fmt.Printf("unexpected type %T\n", x)
	case []string:
		fmt.Printf("[]string %v\n", x)
	case []interface{}:
		fmt.Printf("[]interface{} %v\n", x)
	}
}

func TestReflect(t *testing.T) {
	var o interface{}
	ss := []string{"s", "s"}
	o = ss

	tp := reflect.TypeOf(o)
	fmt.Println("type: ", tp)
	if tp.Kind() == reflect.Slice {
		fmt.Println("t.Kind() == reflect.Slice")
	}
}

func TestGetFirstValueReflect(t *testing.T) {
	ss := []string{"sa", "sa"}
	mm := []map[interface{}]interface{}{
		map[interface{}]interface{}{"aa": "bb", "cc": "dd", "ff": 1010},
		map[interface{}]interface{}{"cc": "dd", "ff": 1010, "aa": "bb"},
	}
	opt := []interface{}{
		[]string{"a", "b", "c"},
		[]interface{}{"a", "b", "c"},
	}

	oo := []interface{}{ss, mm, opt}

	for _, o := range oo {
		tp := reflect.TypeOf(o)
		fmt.Println("type: ", tp)
		if tp.Kind() == reflect.Slice {
			fmt.Println("t.Kind() == reflect.Slice")
		}

		sl := reflect.ValueOf(o)
		fmt.Println(sl.Index(0), "==", sl.Index(1), "=", sl.Index(0) == sl.Index(1))
		fmt.Println(sl.Index(0), "==", sl.Index(1), "=",
			reflect.DeepEqual(sl.Index(0), sl.Index(1)))
		fmt.Println("String()", sl.Index(0).String(), "==", sl.Index(1).String(),
			"=", sl.Index(0).String() == sl.Index(1).String())
		fmt.Println()
	}

}
