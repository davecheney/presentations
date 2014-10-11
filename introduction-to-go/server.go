package main

import "net"

// START OMIT
func server(listener net.Listener) {
	for {
		client, _ := listener.Accept()
		go handle(client)
	}
}

func handle(client net.Conn) {
	// handle connection
}

func main() {
	listener, _ := net.Listen("tcp", ":2000")
	go server(listener)

	// do some other stuff
}

// END OMIT
