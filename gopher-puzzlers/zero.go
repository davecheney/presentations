package main

type Hash [32]byte

func MustNotBeZero(h Hash) {
	if h == Hash{} {
		panic("0")
	}
}

func main() {}
