package main

import (
	"fmt"
	"unicode"
)

// START OMIT
func main() {
	for _, r := range []rune{'ಠ', 'す'} {
		fmt.Printf("%c: letter: %v, lowercase: %v, uppercase: %v\n", r,
			unicode.IsLetter(r), unicode.IsLower(r), unicode.IsUpper(r))
	}
}

// END OMIT
