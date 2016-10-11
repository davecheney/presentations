package main

import "log"
import "sync"

func main() {
	var mu sync.Mutex
	done := make(chan struct{})
	defer func() {
		if e := recover(); e != nil {
			log.Printf("recovered %v", e)
		}
		<-done
	}()
	mu.Lock()
	go func() {
		mu.Lock()
		close(done)
	}()
	mu.Unlock()
	mu.Unlock()
}
