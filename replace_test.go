package tests

import (
	"html/template"
	"os"
	"regexp"
	"testing"
)

// 1000000	2280 ns/op	2.314s
func BenchmarkTemplate(b *testing.B) {
	tmpl := template.New("tt")
	template.Must(tmpl.Delims("${", "}").Parse("azaza ${ .Name } zazaz. "))
	for i := 0; i < b.N; i++ {
		tmpl.Execute(os.Stdout, p)
	}
}

// 1000000	1250 ns/op	1.282s
func BenchmarkRx(b *testing.B) {
	var rx = regexp.MustCompile(`\$\{([\w]{1,1000})\}`)
	for i := 0; i < b.N; i++ {
		rx.ReplaceAllString(`azaza ${Name} zazaz. `, `Mary`)
	}
}
