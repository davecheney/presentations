package main

import "fmt"
import "time"

// START OMIT
type Worker struct {
	stop chan struct{}
}

func (w *Worker) run() {
	defer fmt.Println("All done")
	for {
		select {
		case <-w.stop:
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Waited 100ms")
		}
	}
}

func main() {
	w := &Worker{stop: make(chan struct{})}
	go w.run()
	time.Sleep(300 * time.Millisecond)
	close(w.stop)
}

// END OMIT
