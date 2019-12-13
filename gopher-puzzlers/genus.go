package main

import "fmt"

const (
	Animal = iota
	Mineral
	Vegetable
)

func genus(t uint) {
	switch t {
	case Animal:
		fmt.Println("Animal")
	case Mineral:
		fmt.Println("Mineral")
	case Vegetable:
		fmt.Println("Vegetable")
	}
}

func main() {
	genus(-Animal)
}
