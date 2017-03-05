package main

type T [256]byte // a large stack allocated type

//go:nosplit
func A(t T) {
	B(t)
}

//go:nosplit
func B(t T) {
	C(t)
}

//go:nosplit
func C(t T) {
	D(t)
}

//go:nosplit
//go:noinline
func D(t T) {}

func main() {
	var t T
	A(t)
}
