package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.TrimRight("abcdefedcba", "abcdef")
	fmt.Printf("%q\n", s)
}
