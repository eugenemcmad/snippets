package tests

import "testing"

var (
	tstStrSlice = []string{"a", "b", "c", "d", "e"}
)

// go test -v xr/snippets -bench ^Benchmark_IFSW_ -run ^$

// 70.7 ns/op
func Benchmark_IFSW_IF(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		for i := range tstStrSlice {
			if tstStrSlice[i] == "a" {
				c++
			} else if tstStrSlice[i] == "b" {
				c++
			} else if tstStrSlice[i] == "c" {
				c--
			} else if tstStrSlice[i] == "d" {
				c--
			} else {
				c++
			}
		}
	}
}

// 61.7 ns/op
func Benchmark_IFSW_SWITCH(b *testing.B) {
	var c int
	for i := 0; i < b.N; i++ {
		for i := range tstStrSlice {

			switch tstStrSlice[i] {
			case "a":
				c++
				break
			case "b":
				c++
				break
			case "c":
				c--
				break
			case "d":
				c--
				break
			default:
				c++
			}
		}
	}
}
