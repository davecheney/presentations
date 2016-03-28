package main

import "fmt"

// START OMIT
func main() {
	b := make([]int, 1024)
	b = append(b, 99)
	fmt.Println(cap(b))
}

// END OMIT
