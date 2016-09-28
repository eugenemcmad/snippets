package tests

import (
	"fmt"
	"reflect"
	"testing"
)

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
