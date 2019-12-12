package main

import "fmt"

func main() {
	p, q := new(int), new(int)
	fmt.Println(*p == *q, p == q)

	r, s := new(struct{}), new(struct{})
	fmt.Println(*r == *s, r == s)
}
