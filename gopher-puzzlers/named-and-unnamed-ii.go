package main

func f() {
	// START OMIT
	type stack []uintptr
	var st stack = make([]uintptr, 20)
	// END OMIT
}
