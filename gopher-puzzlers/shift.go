package main

func a() {
	//START1 OMIT
	x := 2 << -1
	//END1 OMIT
}

func b() {
	//START2 OMIT
	x := 1
	y := 2 << x
	//END2 OMIT
}

func c() {
	// START3 OMIT
	var x = uint(1)
	y := 1.0 << x
	// END3 OMIT
}

func d() {
	// START4 OMIT
	x := uint(0)
	y := make([]int, 10)
	z := y[1.0<<x]
	// END4 OMIT
}

// START5 OMIT
func f(int64)

func g(interface{})

func main() {
	var s = uint(0)
	f(1.0 << s) // OK // HL
	g(1.0 << s) // Not OK // HL

	y := make([]int, 10)
	z := y[1.0<<x] // ¯\_(ಠ_ಠ)_/¯ // HL
}

// END5 OMIT
