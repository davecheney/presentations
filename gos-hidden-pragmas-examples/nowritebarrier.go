package main

var count int

//go:nowritebarrier
func increment() {
	count++
}

func main() {
	for i := 0; i < 100; i++ {
		increment()
	}
}
