package main

// START OMIT
type T []int

func F(t T) {}

func main() {
	var x []int
	var y T

	y = x // HL
	F(x)  // HL

	_ = y // keep the compiler happy
}

// END OMIT
