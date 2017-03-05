package main

var x string

//go:noescape
func foo(s string) {
	x = s
}

func main() {
	s := "hello world"
	foo(s)
}
