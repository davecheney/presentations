package main

import "play.ground/complex"

func main() {
	complex.Hello()
}
-- go.mod --
module play.ground
-- complex/complex.go --
package complex

import "fmt"

func Hello() {
	fmt.Println("hello")
}