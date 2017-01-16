package tests

import (
	"fmt"
	"testing"
	"xr/xutor/utils"
)

var (
	testStrForFind = `abcd ${email} aaaaa ${ip} jjjj ${ok} zzzz ${END}`
)

func TestMyFinder(t *testing.T) {

	fmt.Printf("StrFindAllString( true  ): %v \n", utils.StrFindAllString(testStrForFind, "${", "}", true))
	fmt.Printf("StrFindAllString( false ): %v \n", utils.StrFindAllString(testStrForFind, "${", "}", false))

	fmt.Printf("StrFindReplInContent( true  ): %v \n", StrFindReplInContent(testStrForFind, true))
	fmt.Printf("StrFindReplInContent( false ): %v \n", StrFindReplInContent(testStrForFind, false))

}

// 3321 ns/op
func Benchmark_StrFind_RxFind(b *testing.B) {
	z := []byte(testStrForFind)
	for i := 0; i < b.N; i++ {
		testReplRx.FindAll(z, -1)
	}
}

// 1159 ns/op
func Benchmark_StrFind_StrFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.StrFindAllString(testStrForFind, "${", "}", true)
	}
}

// 1165 ns/op
func Benchmark_StrFind_StrFind_NoBorders(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.StrFindAllString(testStrForFind, "${", "}", false)
	}
}


// 1149 ns/op - переменный результат, если и быстрее универсального - то не принципиально
func Benchmark_StrFind_StrFindReplInContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrFindReplInContent(testStrForFind, true)
	}
}

// 1150 ns/op
func Benchmark_StrFind_StrFindReplInContent_NoBorders(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrFindReplInContent(testStrForFind, false)
	}
}


// go test -v xr/snippets -bench ^Benchmark_StrFind_ -run ^$


func StrFindReplInContent(s string, b bool) (res []string) {
	var mox = make(map[int]int)
	var l, na, nb, xl, xr int

	for i := 0; i < len(s); i++ {

		if i+2 > len(s)-1 {
			na = len(s)
		} else {
			na = i + 2
		}

		if i+1 > len(s)-1 {
			nb = len(s)
		} else {
			nb = i + 1
		}

		if s[i:na] == "${" {
			l = i
			i+=1
			continue
		}

		if s[i:nb] == "}" {
			mox[l] = i
			continue
		}
	}

	res = make([]string, len(mox))
	i := 0
	if !b {
		xl = 2
		xr = 1
	}

	for k, v := range mox {
		res[i] = s[k+xl : v+1-xr]
		i++
	}

	return
}