package main

import (
	"fmt"
)

const (
	c1 int = iota
	c2
	c3
	c4 rune = iota
	c5
)

func main() {
	fmt.Println(c5)
}
