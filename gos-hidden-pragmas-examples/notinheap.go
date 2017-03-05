package main

import "fmt"

//go:notinheap
type T struct {
	a int
}

//go:noinline
func F() *T {
	var t T
	return &t
}

func main() {
	fmt.Println(F().a)
}
