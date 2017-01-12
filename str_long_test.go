package tests

import (
	"bytes"
	"strings"
	"testing"
)

var (
	testLongStr = ` azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz. zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz
	                azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz. zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz
	                azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz. zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz
	                azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz.  azaza ${Name} zazaz. zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz`
)

// 10000000	235 ns/op
func Benchmark_StrLongRepl_StringsReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := `${` + k + `}`
		strings.Replace(testLongStr, k2, `Mary`, -1)
	}
}

// 10000000	175 ns/op
func Benchmark_StrLongRepl_StringsPrepReplace(b *testing.B) {
	k := `Name`
	k2 := `${` + k + `}`
	for i := 0; i < b.N; i++ {
		strings.Replace(testLongStr, k2, `Mary`, -1)
	}
}

// 10000000	150 ns/op
func Benchmark_StrLongRepl_BytesReplace(b *testing.B) {
	k := `Name`
	for i := 0; i < b.N; i++ {
		k2 := []byte(`${` + k + `}`)
		bytes.Replace([]byte(testLongStr), k2, []byte(`Mary`), -1)
	}
}

// 10000000	144 ns/op
func Benchmark_StrLongRepl_BytesPrepReplace(b *testing.B) {
	z := []byte(testLongStr)
	k := `Name`
	k2 := []byte(`${` + k + `}`)
	v := []byte(`Mary`)
	for i := 0; i < b.N; i++ {
		bytes.Replace(z, k2, v, -1)
	}
}
