package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	const Second = uint64(time.Second)

	when := -1 * 5 * Second

	fmt.Printf("this happened %v ago\n", when)
}

// END OMIT
