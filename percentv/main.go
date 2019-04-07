package main

import (
	"fmt"
)

// tag::main[]
func main() {
	type T struct {
		I int
	}
	x := []*T{{1}, {2}, {3}}
	y := []*T{{1}, {2}, {4}}

	fmt.Printf("%v %v\n", x, y)
	fmt.Printf("%#v %#v\n", x, y)
}

// end::main[]
