package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	t := time.NewTicker(100)
	t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("tick")
		}
	}
}

// END OMIT
