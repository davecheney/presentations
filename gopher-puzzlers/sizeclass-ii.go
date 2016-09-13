package main

import "fmt"

// START OMIT
func main() {
	b := make([]int, 32)
	b = append(b, 99)
	fmt.Println("len:", len(b), "cap:", cap(b))
}

// END OMIT
