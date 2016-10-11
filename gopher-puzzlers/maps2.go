package main

import "fmt"

// START OMIT
func main() {
	m := make(map[string]int)
	v := m["foo"]
	v++
	m["foo"] = v
	fmt.Println(m["foo"])
}

// END OMIT
