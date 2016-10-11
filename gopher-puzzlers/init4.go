package main

import "fmt"

var x int

func init() {
	x++
}

func main() {
	init()
	fmt.Println(x)
}
