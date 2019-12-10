package main

import "fmt"

func main() {
	var x uint32 = 100.5 + 100.5 + 2147483447.0
	fmt.Println(x)
}
