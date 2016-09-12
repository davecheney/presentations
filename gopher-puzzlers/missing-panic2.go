package main

import "os"

func main() {
	os.Stderr.Close()
	panic("all is lost")
}
