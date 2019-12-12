package main

import "fmt"

var (
	p, q = new(int), new(int)
	r, s = new(struct{}), new(struct{})
)

func main() {
	fmt.Println(*p == *q, p == q)
	fmt.Println(*r == *s, r == s)
}
