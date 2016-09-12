package main

import (
	"fmt"
	"unsafe"
)

// START OMIT
func main() {
	type T struct {
		_ struct{}
		a int
	}
	var t T
	fmt.Println(unsafe.Sizeof(t))
}

//END OMIT
