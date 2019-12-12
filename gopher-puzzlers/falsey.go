package main

import (
	"fmt"
)

func main() {
	var err error
	_, true := err.(error)
	fmt.Println(true)
}
