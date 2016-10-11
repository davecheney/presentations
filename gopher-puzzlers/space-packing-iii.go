package main

func main() {
	type T struct {
		a struct{}
		b int
		c int
		d struct{}
	}

	var t T
	println(&t.a)
	println(&t.b)
	println(&t.c)
	println(&t.d)
}
