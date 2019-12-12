package main

import "fmt"

var x = 1

func a() int {
	x *= 2
	return x
}

func b() int {
	x /= 2
	return x
}

func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(sub(a(), b()))
}
