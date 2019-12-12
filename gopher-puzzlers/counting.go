package main

import (
	"fmt"
)

func main() {
	i := []int{
		1, 2, 3,
	}
	for j := range i {
		fmt.Println(j)
	}
}
