package p

type Work int

// START OMIT
const LIMIT = 10

var semaphore = make(chan struct{}, LIMIT)

func doWork(work *Work) {
	semaphore <- struct{}{} // acquire semaphore
	// do work
	<-semaphore // release semaphore
}

// END OMIT
