package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[:]   // slice of all the elements
	c := a[3:]  // slice of the 4th element through to the end
	d := a[:6]  // slice of the first 6 elements
	e := a[3:6] // slice of the 4th, 5th, 6th elements
	fmt.Printf("a - %v\n", a)
	fmt.Printf("b - %v\n", b)
	fmt.Printf("c - %v\n", c)
	fmt.Printf("d - %v\n", d)
	fmt.Printf("e - %v\n", e)

	// y := [i:j] --> i'th index is inclusive and j'th index is  exclusive i.e (inclusive til j-1)

	slice := make([]int, 3, 100)
	fmt.Printf("slice - %v\n", slice)
	fmt.Printf("Length - %v\n", len(slice))
	fmt.Printf("Capacity - %v\n", cap(slice))

	x := []int{}
	fmt.Printf("x - %v\n", x)
	fmt.Printf("Length - %v\n", len(x))
	fmt.Printf("Capacity - %v\n", cap(x))

	x = append(x, 1)

	fmt.Printf("x - %v\n", x)
	fmt.Printf("Length - %v\n", len(x))
	fmt.Printf("Capacity - %v\n", cap(x))

	x = append(x, []int{2, 3, 4, 5}...) // equivalent to -> x = append(x, 2, 3, 4, 5)

	fmt.Printf("x - %v\n", x)
	fmt.Printf("Length - %v\n", len(x))
	fmt.Printf("Capacity - %v\n", cap(x))
}
