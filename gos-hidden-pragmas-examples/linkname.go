package main

import "fmt"
import _ "unsafe"

// bytes.Equal is implemented in runtime/asm_$goarch.s
//go:linkname bytesEqual bytes.Equal
func bytesEqual(x, y []byte) bool

func main() {
	fmt.Println(bytesEqual(make([]byte, 1), make([]byte, 2)))
}
