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
	switch t := o.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case []string:
		fmt.Printf("[]string %v\n", t)
	case []interface{}:
		fmt.Printf("[]interface{} %v\n", t)
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
