package main

import "fmt"

func main() {

	// map[key]value -- hashmap equivalent

	mapExample := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  6,
	}

	fmt.Println(mapExample)

	// creating maps using make function

	mapAnotherExample := make(map[string]int)

	mapAnotherExample = map[string]int{
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
	}

	mapAnotherExample["eleven"] = 11

	fmt.Printf("%d", mapAnotherExample["eleven"])

	fmt.Println(mapAnotherExample)

	delete(mapAnotherExample, "eleven")

	fmt.Println(mapAnotherExample)

	// fmt.Println(mapAnotherExample["eleven"], ok)

	fmt.Println(len(mapAnotherExample))

	// checking if a key is present in the map
	temp, ok := mapAnotherExample["eleven"]
	fmt.Println(temp, ok)

	tem := mapAnotherExample
	delete(tem, "ten")
	fmt.Println(tem)
	fmt.Println(mapAnotherExample)
}
