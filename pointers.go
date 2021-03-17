package main

import "fmt"

type Sample struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a++
	fmt.Println(a, *b)
	*b = 50
	fmt.Println(a, *b)

	x := [3]int{1, 2, 3}
	y := &x[1]
	z := &x[2]

	// %v - value -- %p - pointer
	fmt.Printf("%v %p %p\n", x, y, z)

	var sample *Sample
	var sample1 *Sample

	fmt.Printf("sample1 - %v", sample1)

	sample = &Sample{
		foo: 42,
	}
	sample1 = new(Sample)
	(*sample1).foo = 32

	fmt.Println(sample)
	fmt.Println(sample1)
}
