package main

import (
	"fmt"
	"time"
)

// START OMIT
func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func main() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}

// END OMIT
