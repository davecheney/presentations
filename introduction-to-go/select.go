package main

import "time"
import "fmt"

// START OMIT
func SendWithTimeout(data string, to chan int, timeout time.Duration) error {
	select {
	case to <- data:
		return nil // everything worked
	case <-time.After(timeout):
		return fmt.Errorf("timeout after %s", timeout)
	}
}

// END OMIT
func main() {
}
