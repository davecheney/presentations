package main

import "fmt"

func main() {
	// START OMIT
	q := "Tony Abbott"
	p := &q

	*p = "Kevin Rudd"
	fmt.Println(q)
	// END OMIT
}
