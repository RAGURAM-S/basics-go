package main

import (
	"fmt"
	"math"
)

const a int16 = 27

// iota is scoped to a constant block
const (
	p = iota
	q
	r
)

const (
	// default value of iota is 0 and increments within the block scope
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {

	// typed constants

	// constants should be assignable during compile time
	const myConstant int = 2 * 24
	fmt.Printf("%v - %T\n", myConstant, myConstant)

	const pi float64 = math.Pi
	fmt.Printf("%v- %T\n", pi, pi)

	// shadowing can be done with constants -- inner scope has greater precedence
	fmt.Printf("a -> %v- %T\n", a, a)

	const a int = 12
	const b string = "hello"
	const c bool = true
	const d float32 = 12.24

	fmt.Printf("a -> %v- %T\n", a, a)
	fmt.Printf("b -> %v- %T\n", b, b)
	fmt.Printf("c -> %v- %T\n", c, c)
	fmt.Printf("d -> %v- %T\n", d, d)

	// untyped constants

	const x = 75
	fmt.Printf("x -> %v- %T\n", x, x)
	var y int16 = 10
	fmt.Printf("x+y -> %v- %T\n", x+y, x+y)

	// enumerated constants
	fmt.Printf("p -> %v - %T\n", p, p)
	fmt.Printf("q -> %v - %T\n", q, q)
	fmt.Printf("r -> %v - %T\n", r, r)

	//
	fileSize := 4000000000.0
	fmt.Printf("%.2f MB", fileSize/GB)

}
