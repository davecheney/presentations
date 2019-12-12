package main

import "fmt"

const (
	Animal = iota
	Mineral
	Vegitable
)

func genus(t uint) {
	switch t {
	case Animal:
		fmt.Println("Animal")
	case Mineral:
		fmt.Println("Mineral")
	case Vegitable:
		fmt.Println("Vegitable")
	}
}

func main() {
	genus(-Animal)
}
