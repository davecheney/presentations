package main

import "fmt"

var x int

func init() {
	main()
}

func main() {
	x++
	fmt.Println(x)
}
