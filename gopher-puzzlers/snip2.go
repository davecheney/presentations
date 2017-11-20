package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.TrimSuffix("abcdefedcba", "abcdef")
	fmt.Printf("%q\n", s)
}
