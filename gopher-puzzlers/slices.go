package main

import "fmt"

func main() {
	b := []byte("Hello")
	r := []rune("Seattle")
	i := []int("Gophers")
	fmt.Println(b, r, i)
}
