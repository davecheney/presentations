package main

import "fmt"
import "time"
import "launchpad.net/tomb"

// START OMIT
type Worker struct {
	tomb.Tomb
}

func (w *Worker) run() {
	defer w.Tomb.Done()
	defer fmt.Println("All done")
	for {
		select {
		case <-w.Tomb.Dying():
			return
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Waited 100ms")
		}
	}
}

func main() {
	w := &Worker{}
	go w.run()
	<-time.After(300 * time.Millisecond)
	w.Tomb.Kill(nil) // normal exit
	w.Tomb.Wait()
}

// END OMIT
