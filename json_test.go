package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 500000	      2629 ns/op	     608 B/op	      16 allocs/op
func BenchmarkJson(b *testing.B) {
	b.ReportAllocs()

	m := map[string]int{"a": 10, "b": 20, "c": 30, "d": 40, "e": 50}
	s := ""
	for i := 0; i < b.N; i++ {
		j, err := json.Marshal(m)
		if err != nil {
			b.FailNow()
		}
		s = string(j)
	}
	fmt.Println(s)
}

// 500000	      3734 ns/op	     528 B/op	      28 allocs/op
func BenchmarkSprintf(b *testing.B) {
	b.ReportAllocs()

	m := map[string]int{"a": 10, "b": 20, "c": 30, "d": 40, "e": 50}
	s := ""
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%#v", m)
	}
	fmt.Println(s)
}
