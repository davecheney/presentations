package main

// START OMIT
type T []int

func F(t T) {}

func main() {
	var q []int
	F(q)
}

// END OMIT
