package main

import "time"
import "fmt"

func main() {
	// START OMIT
	var work []string
	select {
	case w := <-incoming:
		work = append(work, w)
	default:
		// will fire if incoming is not ready
	}
	// END OMIT
}
