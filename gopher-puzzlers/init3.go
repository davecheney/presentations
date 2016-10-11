package main

import "fmt"
import "runtime"

func init() {
	var pcs [1]uintptr
	runtime.Callers(1, pcs[:])
	fn := runtime.FuncForPC(pcs[0])
	fmt.Println(fn.Name())
}

func main() {}
