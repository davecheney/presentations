package main

import "fmt"

//go:noescape
func length(s string) int

func main() {
	s := "hello world"
	l := length(s)
	fmt.Println(l)
}
