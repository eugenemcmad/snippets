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

}

// go test -v xr/snippets -bench ^Benchmark_StrFind_ -run ^$

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






