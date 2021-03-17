package main

import "fmt"

func main() {

	// array takes as many values in the
	scores := [...]int{85, 87, 90, 75}
	fmt.Printf("Scores %v  -- %T\n", scores, scores)

	var students [3]string
	students[0] = "Bruce Wayne"
	students[1] = "Tony Stark"
	students[2] = "Sherlock Holmes"
	fmt.Printf("Students %v\n", students)
	fmt.Printf("No of students %v\n", len(students))

	// 2d arrays

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Printf("Identity Matirx - %v\n", identityMatrix)

	// copying in arrays in go
	// copying of arrays in go lang creates a copy of the array and assigns it to a new variable

	a := [...]int{1, 2, 3}
	b := a // a is an array and b is another array in itself
	b[1] = 20
	fmt.Printf("a -> %v\n", a)
	fmt.Printf("b -> %v\n", b)

	// to overcome this issue, we use pointers

	x := [...]int{1, 2, 3}
	y := &x // x is the array and y is pointing to x
	y[1] = 20
	fmt.Printf("x -> %v\n", x)
	fmt.Printf("y -> %v\n", y)

	// slices

	q := []int{9, 8, 7}
	r := q
	fmt.Printf("q - %v %T\n", q, q)
	fmt.Printf("Length %v \nCapacity %v\n", len(q), cap(q))
	r[0] = 1
	fmt.Printf("q - %v\n", q)
	fmt.Printf("r - %v\n", r)
}
