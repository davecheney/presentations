package main

import (
	"fmt"
	"unsafe"
)

// START OMIT
func main() {
	const n = 4 >> (^uint(0) >> 63) // HL
	type X [n]uint
	type Y [n]int

	var x X
	var y Y
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(y))
}

// END OMIT
