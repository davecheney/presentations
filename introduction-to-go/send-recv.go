package main

import "fmt"

// START OMIT
func greeting(c chan string) {
	words := []string{"Welcome", "to", "Go"}
	for _, word := range words {
		c <- word
	}
	close(c)
}

func main() {
	c := make(chan string)
	go greeting(c)

	for word := range c {
		fmt.Println(word)
	}
}

// END OMIT
