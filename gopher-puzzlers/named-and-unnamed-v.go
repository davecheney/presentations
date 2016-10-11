package main

// START OMIT
type T []int

func F(t T) {}

func main() {
	var x []int
	var y T

	y = x // HL
	F(x)  // HL
}

// END OMIT
