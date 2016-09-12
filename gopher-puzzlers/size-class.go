package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 1024)
	b = append(b, 42)

	fmt.Println(len(b), cap(b))
}
