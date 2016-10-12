package main

import "fmt"

func main() {
	go func() { panic("Boom") }()
	for i := 0; i < 10000000; {
		fmt.Print(".")
	}
	fmt.Println("Done")
}
