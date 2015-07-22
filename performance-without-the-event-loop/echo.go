package main

import (
	"io"
	"net"
	"log"
)

// START OMIT
func echo(rw io.ReadWriteCloser) {
	defer rw.Close()
	io.Copy(rw, rw)
}

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("could not accept connection: %v", err)
		}
		go echo(c)
	}
}
// END OMIT
