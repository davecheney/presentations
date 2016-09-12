package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}

// END OMIT
