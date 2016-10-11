package main

import "fmt"

// START OMIT

func main() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- 1:
		case ch <- 2:
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}

// END OMIT
