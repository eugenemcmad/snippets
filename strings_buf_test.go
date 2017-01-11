package tests

import (
	"testing"

	"bytes"
	"fmt"
	"strconv"
)

// 6150 ns/op
func Benchmark_SB_Sprintf(b *testing.B) {
	var srcs = getTestStrSlice()

	for i := 0; i < b.N; i++ {
		var res string
		for j, _ := range srcs {
			res = fmt.Sprintf("%s %s", res, srcs[j])
		}
	}
}

// 3111 ns/op
func Benchmark_SB_StrPlusEq(b *testing.B) {
	var srcs = getTestStrSlice()

	for i := 0; i < b.N; i++ {
		var res string
		for j, _ := range srcs {
			res += srcs[j]
			res += " "
		}
	}
}

// 928 ns/op
func Benchmark_SB_BufWrStr(b *testing.B) {
	var srcs = getTestStrSlice()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j, _ := range srcs {
			buf.WriteString(srcs[j])
			buf.WriteString(" ")
		}
		buf.String()
	}
}

// 928 ns/op
func Benchmark_SB_BufWr(b *testing.B) {
	var bsrcs = getTestBytesSlice()
	var z = []byte(" ")

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j, _ := range bsrcs {
			buf.Write(bsrcs[j])
			buf.Write(z)
		}
		buf.String()
	}
}

func getTestStrSlice() (srcs []string) {
	for i := 10000; i < 10020; i++ {
		srcs = append(srcs, strconv.Itoa(i))
	}
	return
}

func getTestBytesSlice() (bsrcs [][]byte) {
	var tmp = getTestStrSlice()
	for j, _ := range tmp {
		bsrcs = append(bsrcs, []byte(tmp[j]))
	}
	return
}
