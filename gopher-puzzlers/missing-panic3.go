package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stderr.Close()
	f, _ := os.Create("/tmp/wut")
	fmt.Println(f.Fd())
	defer f.Close()
	panic("all is lost")
}
