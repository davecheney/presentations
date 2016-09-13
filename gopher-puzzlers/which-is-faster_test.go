package makebench

import "testing"

var result []byte

// START OMIT
func BenchmarkOneLiteralByte(b *testing.B) {
	var v []byte
	for i := 0; i < b.N; i++ {
		v = []byte{0}
	}
	result = v
}

func BenchmarkMakeOneByte(b *testing.B) {
	var v []byte
	for i := 0; i < b.N; i++ {
		v = make([]byte, 1)
	}
	result = v
}

// END OMIT
