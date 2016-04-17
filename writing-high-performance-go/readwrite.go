package main

import "fmt"

// START OMIT
// sendfile sends the contents of path to the client c.
func sendfile(c net.Conn, path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	// Set the deadline to one minute from now.
	c.SetWriteDeadline(time.Now().Add(60 * time.Second))

	// Copy will send as much of r to the client as it can in 60 seconds.
	_, err = io.Copy(c, r)
	return err
}

// END OMIT
