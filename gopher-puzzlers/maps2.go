package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int)
	v := m["foo"] // HL
	v++           // HL
	m["foo"] = v  // HL
	fmt.Println(m["foo"])
}

// END OMIT
