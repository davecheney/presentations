package main

import "fmt"

func f(a int, b uint) {
	var min = 0
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func main() {
	f(9000, 314)
}
