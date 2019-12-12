package main

import "fmt"

func main() {
	s := make([]int, 0, 8)
	r := s[2:6:7]
	fmt.Println(len(r), cap(r))
}
