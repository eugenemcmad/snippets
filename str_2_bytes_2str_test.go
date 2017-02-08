package tests

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

var (
	testLtlHtmlAsStr   = `<!-- very short string -->`
	testLtlHtmlAsBytes = []byte(testLtlHtmlAsStr)

	testBigHtmlAsStr   = `<!-- padding -->${replacement}<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <!--question --> <div align="center"><font face="Tahoma, Arial, Helvetica, sans-serif" size="4" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 23px; color:#126184;background-color:rgba(255,255,255,0.7);"> Which one of the numbers does not belong in the following series? - 2 - 3 - 6 - 7 - 8 - 14 - 15 - 30 <!-- padding --><div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> </span></font> </div> <!--question END--> <!-- padding --><div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <table border="0" cellspacing="0" cellpadding="0" width="90%"> <tr><td align="center" style="line-height:30px;"> <font face="Tahoma, Arial, Helvetica, sans-serif" size="4" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 27px; color:#126184;"> Choose the correct answer: </span></font> </td></tr> </table> <!-- padding --><div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <a href="http://questandrest.com/question/result?question=12&_xxhash=${xxhash64hex16}&answer=1&type=full" target="_blank" style="text-decoration: none;"> <table border="0" cellspacing="0" cellpadding="10" width="56%" style="min-width:300px;"> <tr> <td align="center" style="border: solid; border-color: #126184;"> <font face="Tahoma, Arial, Helvetica, sans-serif" size=4"" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 27px; color:#126184;">THREE</span> </font> </td> </tr> </table> </a> <div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <a href="http://questandrest.com/question/result?question=12&_xxhash=${xxhash64hex16}&answer=2&type=full" target="_blank" style="text-decoration: none;"> <table border="0" cellspacing="0" cellpadding="10" width="56%" style="min-width:300px;"> <tr> <td align="center" style="border: solid; border-color: #126184;"> <font face="Tahoma, Arial, Helvetica, sans-serif" size=4"" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 27px; color:#126184;">FIFTEEN</span> </font> </td> </tr> </table> </a> <div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <a href="http://questandrest.com/question/result?question=12&_xxhash=${xxhash64hex16}&answer=3&type=full" target="_blank" style="text-decoration: none;"> <table border="0" cellspacing="0" cellpadding="10" width="56%" style="min-width:300px;"> <tr> <td align="center" style="border: solid; border-color: #126184;"> <font face="Tahoma, Arial, Helvetica, sans-serif" size=4"" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 27px; color:#126184;">EIGHT</span> </font> </td> </tr> </table> </a> <div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div> <a href="http://questandrest.com/question/result?question=12&_xxhash=${xxhash64hex16}&answer=4&type=full" target="_blank" style="text-decoration: none;"> <table border="0" cellspacing="0" cellpadding="10" width="56%" style="min-width:300px;"> <tr> <td align="center" style="border: solid; border-color: #126184;"> <font face="Tahoma, Arial, Helvetica, sans-serif" size=4"" color="#126184" style="font-size:27px;"> <span style="font-family: Tahoma, Arial, Helvetica, sans-serif; font-size: 27px; color:#126184;">SEVEN</span> </font> </td> </tr> </table> </a> <!-- padding --><div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div>`
	testBigHtmlAsBytes = []byte(testBigHtmlAsStr)

	testReplStr   = `${replacement}`
	testReplBytes = []byte(testReplStr)
)

// go test -v xr/snippets -bench ^Benchmark_SBS_ -run ^$

func Benchmark_SBS_LtLBytesToStringMutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BytesToString(testLtlHtmlAsBytes)
	}
}

func Benchmark_SBS_LtlBytesToStringImmutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testLtlHtmlAsBytes)
	}
}

func Benchmark_SBS_LtlStringToBytesMutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringToBytesUnsafe(testLtlHtmlAsStr)
	}
}

func Benchmark_SBS_LtlStringToBytesImmutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testLtlHtmlAsStr)
	}
}

func Benchmark_SBS_BigBytesToStringMutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BytesToString(testBigHtmlAsBytes)
	}
}

func Benchmark_SBS_BigBytesToStringImmutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBigHtmlAsBytes)
	}
}

func Benchmark_SBS_BigStringToBytesMutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringToBytesUnsafe(testBigHtmlAsStr)
	}
}

func Benchmark_SBS_BigStringToBytesImmutable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testBigHtmlAsStr)
	}
}

func TestStringToBytes(t *testing.T) {
	b1 := StringToBytesUnsafe(testBigHtmlAsStr)
	fmt.Println(b1)
	b2 := []byte(testBigHtmlAsStr)
	fmt.Println(b2)
	if len(b1) != len(b2) {
		t.Errorf("len %d != %d\n", len(b1), len(b2))
	}
	if cap(b1) != cap(b2) {
		fmt.Printf("cap %d != %d\n", cap(b1), cap(b2))
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			t.Fail()
		}
	}

	// bytes.Contains(b1, testReplBytes) // ERROR! По полученым так байтам нельзя ходить, применять к ним операции!

	if !bytes.Contains(testBigHtmlAsBytes, StringToBytesUnsafe(testReplStr)) {
		t.Errorf("'%s' not found\n", testReplStr)
	}

	var buf = bytes.NewBuffer(b2)
	buf.Write(b1)
	if bytes.Count(buf.Bytes(), testReplBytes) != 2 {
		t.Errorf("'%s' wrong count\n", testReplBytes)
	}
}

func TestBytesToString(t *testing.T) {
	s1 := BytesToString(testBigHtmlAsBytes)
	s2 := string(testBigHtmlAsBytes)

	if len(s1) != len(s2) {
		t.Errorf("len %d != %d\n", len(s1), len(s2))
	}

	if s1 != s2 {
		t.Fail()
	}

	strings.Contains(s1, testReplStr) // OK
}

// StringToBytes accepts string and returns their []byte presentation
// instead of byte() this method doesn't generate memory allocations,
// BUT it is not safe to use anywhere because it points
// this helps on 0 memory allocations
// cap(result) == 0
func StringToBytesUnsafe(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
