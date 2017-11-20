package main

import "fmt"

type P *int
type Q *int

func main() {
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	fmt.Println(*p, *q)
}
