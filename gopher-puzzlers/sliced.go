package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3, 8)
	s[2:5][0] = 5
	fmt.Println(s)
}
