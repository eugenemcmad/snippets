package tests

import (
	"bytes"
	"fmt"
	"strconv"
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

func TestBuffer2(t *testing.T) {
	var emln = "test@aol.com"
	var eid = 555
	var rs int64 = 1478601200

	var buffer bytes.Buffer
	buffer.WriteString(`{"email":"`)
	buffer.WriteString(emln)
	buffer.WriteString(`","eid":"`)
	buffer.WriteString(strconv.Itoa(eid))
	buffer.WriteString(`","ers":"`)
	buffer.WriteString(strconv.FormatInt(rs, 10))
	buffer.WriteString(`"}`)

	leadStr := buffer.String()

	fmt.Println(leadStr)
}
