package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const n = 4
	type X [n]uint
	type Y [n]int

	var x X
	var y Y
	fmt.Println(unsafe.Sizeof(x), unsafe.Sizeof(y))
}
