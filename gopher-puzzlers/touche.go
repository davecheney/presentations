package main

import "fmt"

func main() {
	s := string("touché")
	b := []byte{'t', 'o', 'u', 'c', 'h', 'é'}
	fmt.Println(len(s) == len(b))
}
