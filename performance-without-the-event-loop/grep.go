package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func grep(r io.Reader, needle string) {
	br := bufio.NewReader(r)
	lines := make(chan string, 20)

	go func() {
		defer close(lines)
		for {
			line, err := br.ReadString('\n')
			if err != nil {
				return
			}
			lines <- line
		}
	}()

	for line := range lines {
		if strings.Contains(line, needle) {
			fmt.Println(line)
		}
	}
}

func main() {
	grep(nil, "")
}
