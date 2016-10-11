// START OMIT
package main

func A(string string) string {
	return string + string
}

func B(len int) int {
	return len+len
}

func C(val, default string) string {
	if val == "" {
		return default
	}
	return val
}

// END OMIT

func main() {}
