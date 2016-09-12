package main

import (
	"fmt"
	"unsafe"
)

// START OMIT
func main() {
	type T struct {
		a int
		_ struct{}
	}
	var t T
	fmt.Println(unsafe.Sizeof(t))
}

//END OMIT
