package main

import "fmt"
import "time"
import "sync"

// START OMIT
type Worker struct {
	stop chan struct{}
	wg   sync.WaitGroup
}

func (w *Worker) run() {
	defer w.wg.Done()
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
	w.wg.Add(1)
	go w.run()
	time.Sleep(300 * time.Millisecond)
	close(w.stop)
	w.wg.Wait()
}

// END OMIT
