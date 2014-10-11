package main

import "fmt"

// START1 OMIT
type Square struct{ length int }

func (s *Square) Height() int { return s.length }
func (s *Square) Width() int  { return s.length }

type Rectangle struct{ height, width int }

func (r *Rectangle) Height() int { return r.height }
func (r *Rectangle) Width() int  { return r.width }

// END1 OMIT

// START2 OMIT
type Parallelogram interface {
	Height() int
	Width() int
}

func Area(p Parallelogram) {
	fmt.Printf("Height: %d, Width: %d, Area: %d\n",
		p.Height(), p.Width(), p.Height()*p.Width())
}

func main() {
	s := &Square{length: 10}
	r := &Rectangle{height: 4, width: 25}

	Area(s)
	Area(r)
}

// END2 OMIT
