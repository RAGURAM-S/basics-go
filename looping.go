package main

import "fmt"

func main() {

	for i, j := 0, 1; i <= 10; i, j = i+2, j+2 {
		fmt.Println(i, j)
	}

	fmt.Println("while loops")
	// while loops
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("infinite loops")
	// infinite loops
	y := 2
	for {
		if y >= 7 {
			break
		}
		fmt.Println(y)
		y++
	}

	fmt.Println("nested loops")
	// nested loops
	for a := 1; a <= 5; a++ {
		for b := 1; b <= 5; b++ {
			fmt.Printf("%v * %v = %v\n", a, b, a*b)
		}
	}

	// breaks
	fmt.Println("break statement with labels")
Loop:
	for a := 1; a <= 5; a++ {
		for b := 1; b <= 5; b++ {
			fmt.Printf("%v * %v = %v\n", a, b, a*b)
			if a*b >= 5 {
				break Loop
			}
		}
	}

	// looping through collections

	fmt.Println("looping through slices/ arrays")

	slice := []int{1, 2, 3, 4, 5}
	for key, value := range slice {
		fmt.Println(key, value)
	}

	fmt.Println("looping through maps")

	mapExample := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  6,
	}

	for key, value := range mapExample {
		fmt.Println(key, value)
	}

	fmt.Println("looping through string")
	str := "hello"
	for _, value := range str {
		fmt.Println(string(value))
	}
}
