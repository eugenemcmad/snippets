package tests

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	var buf bytes.Buffer
	var pars = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	for k, v := range pars {
		buf.WriteString("&")
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)
	}

	fmt.Printf("%+v\n", buf)

	fmt.Println(buf.Len())

	fmt.Println(buf.Bytes())

	fmt.Println(buf.String())

	fmt.Printf("%+v\n", buf)

	fmt.Println(buf.Len())

	fmt.Println(buf.Bytes())

	fmt.Println(buf.String())
}
