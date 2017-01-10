package tests

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

var (
	p = tPers{Name: "Mary"}

	fmap = template.FuncMap{
		"gmap": gmap,
	}
)

type tPers struct {
	Name string //exported field since it begins with a capital letter
}

func TestTemplate1(t *testing.T) {
	var err error
	tmpl := template.New("tt")
	x := template.Must(tmpl.Delims("${", "}").Parse("azaza ${ .Name } zazaz. "))
	tmpl.Execute(os.Stdout, p)
	fmt.Printf("\n %v (%v) \n", *x, err)
}

func TestTemplate2(t *testing.T) {
	var err error
	tmpl := template.New("tt")
	x := template.Must(tmpl.Delims("${", "}").Parse("azaza ${ . } zazaz. "))
	tmpl.Execute(os.Stdout, "Gary")
	fmt.Printf("\n %v (%v) \n", *x, err)
}

func TestTemplate3(t *testing.T) {
	var err error
	tmpl := template.New("tt")
	x := template.Must(tmpl.Delims("${", "}").Funcs(fmap).Parse(`azaza ${ gmap . "g" } zazaz. `))
	tmpl.Execute(os.Stdout, map[string]string{"g": "Gary", "l": "Lary"})
	fmt.Printf("\n %v (%v) \n", *x, err)
}

func gmap(m map[string]string, k string) string {
	return m[k]
}
