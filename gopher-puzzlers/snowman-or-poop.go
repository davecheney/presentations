package main

import (
	"log"
	"os"
	"time"
)

// START OMIT
func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		r.Close()
	}()

	var buf [16]byte
	_, err = r.Read(buf[:])
	if err != nil {
		log.Println("☃")
	} else {
		log.Println("💩")
	}

	w.Close()
}

// END OMIT
