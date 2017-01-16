package tests

import (
	"bytes"
	"html/template"
	"strings"
	"testing"

	"github.com/alexflint/go-restructure/regex"
	"github.com/dlclark/regexp2"
)

var (
	testReplTmpl = template.New("tt")
)

// go test -v xr/snippets -bench ^Benchmark_Repl_ -run ^$

//// 1000000	7579 ns/op
//func Benchmark_Repl_TemplateReplace(b *testing.B) {
//	template.Must(testReplTmpl.Delims("${", "}").Parse("azaza ${ .Name } zazaz. "))
//	for i := 0; i < b.N; i++ {
//		testReplTmpl.Execute(os.Stdout, p)
//	}
//}

// 1000000	1270 ns/op
func Benchmark_Repl_RxReplaceAllString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testReplRx.ReplaceAllString(`azaza ${Name} zazaz. `, `Mary`)
	}
}

// 1000000	1246 ns/op
func Benchmark_Repl_RxReplaceAll(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		testReplRx.ReplaceAll(z, v)
	}
}

// 10000000	242 ns/op
func Benchmark_Repl_StringsReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := `${` + k + `}`
		strings.Replace(`azaza ${Name} zazaz. `, k2, `Mary`, -1)
	}
}

// 10000000	175 ns/op
func Benchmark_Repl_StringsPrepReplace(b *testing.B) {
	k := `Name`
	k2 := `${` + k + `}`
	for i := 0; i < b.N; i++ {
		strings.Replace(`azaza ${Name} zazaz. `, k2, `Mary`, -1)
	}
}

// 10000000	156 ns/op
func Benchmark_Repl_BytesReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := []byte(`${` + k + `}`)
		bytes.Replace([]byte(`azaza ${Name} zazaz. `), k2, []byte(`Mary`), -1)
	}
}

// 10000000	144 ns/op
func Benchmark_Repl_BytesPrepReplace(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	k := `Name`
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		k2 := []byte(`${` + k + `}`)
		bytes.Replace(z, k2, v, -1)
	}
}

// 10000000	144 ns/op
func Benchmark_Repl_BytesPrep2Replace(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	k := `Name`
	k2 := []byte(`${` + k + `}`)
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		bytes.Replace(z, k2, v, -1)
	}
}

// 1000000	1234 ns/op
// alexflint/go-restructure/regex
func Benchmark_Repl_RxRestReplaceAllString(b *testing.B) {
	z := `azaza ${Name} zazaz. `
	rx := regex.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAllString(z, `Mary`)
	}
}

// 1000000	1226 ns/op
// alexflint/go-restructure/regex
func Benchmark_Repl_RxRestReplaceAll(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	v := []byte(`Mary`)
	rx := regex.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAll(z, v)
	}
}

// 500000	2832 ns/op
// dlclark/regexp2
func Benchmark_Repl_Rx2Replace(b *testing.B) {
	z := `azaza ${Name} zazaz. `
	rx := regexp2.MustCompile(`\$\{([\w]{1,1000})\}`, 0)
	for i := 0; i < b.N; i++ {
		rx.Replace(z, `Mary`, -1, -1)
	}
}
