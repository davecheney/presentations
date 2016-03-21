package copy

import (
	"testing"
)

// START OMIT
var c = []byte("1789678900001234567890")

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := make([]byte, len(c))
		copy(d, c)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = append([]byte(nil), c...)
	}
}

// END OMIT
