package main

var v int

//go:norace
func add() {
	v++
}

func main() {
	for i := 0; i < 5; i++ {
		go add()
	}
}
