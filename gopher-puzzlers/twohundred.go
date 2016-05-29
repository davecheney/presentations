package main

import (
	"fmt"
)

// START OMIT
func main() {
	x := []int{100, 200, 300, 400, 500, 600, 700}
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)
}

// END OMIT
