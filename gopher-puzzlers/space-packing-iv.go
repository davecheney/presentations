package main

import (
	"fmt"
	"unsafe"
)

// START OMIT
func main() {
	type T struct {
		a int     // 4 or 8 bytes
		_ [1]byte // 1 byte, padded to 4 or 8
	}
	var t T
	fmt.Println(unsafe.Sizeof(t))
}

//END OMIT
