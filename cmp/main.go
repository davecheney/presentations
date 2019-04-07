package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// tag::main[]
func main() {
	type T struct {
		I int
	}
	x := []*T{{1}, {2}, {3}}
	y := []*T{{1}, {2}, {4}}

	diff := cmp.Diff(x, y)
	fmt.Printf(diff)
}

// end::main[]
