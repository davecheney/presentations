package fib

// STARTBENCH OMIT
import "testing"

func BenchmarkFib(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

// ENDBENCH OMIT

// STARTFIB OMIT
// Fib computes the n'th number in the Fibonacci series.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// ENDFIB OMIT
