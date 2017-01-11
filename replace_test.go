package tests

import (
	"bytes"
	"html/template"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/alexflint/go-restructure/regex"
	"github.com/dlclark/regexp2"
)

// 1000000	2280 ns/op
func Benchmark_Repl_TemplateReplace(b *testing.B) {
	tmpl := template.New("tt")
	template.Must(tmpl.Delims("${", "}").Parse("azaza ${ .Name } zazaz. "))
	for i := 0; i < b.N; i++ {
		tmpl.Execute(os.Stdout, p)
	}
}

// 1000000	1216 ns/op
func Benchmark_Repl_RxReplaceAllString(b *testing.B) {
	var rx = regexp.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAllString(`azaza ${Name} zazaz. `, `Mary`)
	}
}

// 1000000	1206 ns/op
func Benchmark_Repl_RxReplaceAll(b *testing.B) {
	var rx = regexp.MustCompile(`\$\{([\w]{1,1000})\}`)
	z := []byte(`azaza ${Name} zazaz. `)
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAll(z, v)
	}
}

// 10000000	235 ns/op
func Benchmark_Repl_StringsReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := `${` + k + `}`
		strings.Replace(`azaza ${Name} zazaz. `, k2, `Mary`, -1)
	}
}

// 10000000	150 ns/op
func Benchmark_Repl_BytesReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := []byte(`${` + k + `}`)
		bytes.Replace([]byte(`azaza ${Name} zazaz. `), k2, []byte(`Mary`), -1)
	}
}

// 10000000	135 ns/op
func Benchmark_Repl_BytesPreReplace(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	k := `Name`
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		k2 := []byte(`${` + k + `}`)
		bytes.Replace(z, k2, v, -1)
	}
}

// 1000000	1218 ns/op
// alexflint/go-restructure/regex
func Benchmark_Repl_RxRestReplaceAllString(b *testing.B) {
	z := `azaza ${Name} zazaz. `
	rx := regex.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAllString(z, `Mary`)
	}
}

// 1000000	1210 ns/op
// alexflint/go-restructure/regex
func Benchmark_Repl_RxRestReplaceAll(b *testing.B) {
	z := []byte(`azaza ${Name} zazaz. `)
	v := []byte(`Mary`)
	rx := regex.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAll(z, v)
	}
}

// 500000	2799 ns/op
// dlclark/regexp2
func Benchmark_Repl_Rx2Replace(b *testing.B) {
	z := `azaza ${Name} zazaz. `
	rx := regexp2.MustCompile(`\$\{([\w]{1,1000})\}`, 0)
	for i := 0; i < b.N; i++ {
		rx.Replace(z, `Mary`, -1, -1)
	}
}
