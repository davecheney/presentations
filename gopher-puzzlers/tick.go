package main

import "time"

func main() {
	t := time.NewTicker(100)
	for range t.C {
		t.Stop()
	}
}
