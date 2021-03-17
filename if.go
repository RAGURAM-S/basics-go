package main

import "fmt"

func typeSwitch() {
	var i interface{} = "hello"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case float64:
		fmt.Println("i is a float31")
	case string:
		fmt.Println("i is a string")
	case bool:
		fmt.Println("i is a bool")
	default:
		fmt.Println("another data type")
	}
}

func main() {
	mapExample := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  6,
	}

	// variables - pop, ok - are scoped only to the if block and cannot be accessed outside the if block
	if pop, ok := mapExample["one"]; ok {
		fmt.Println(pop, ok)
	}

	if pop, ok := mapExample["seven"]; !ok {
		fmt.Println(pop, ok)
	}

	// golang has implicit breaks and explicit fallthroughs in the switch statments

	switch i := 2 + 3; i {
	case 1, 3, 5:
		fmt.Println("one three or five")
	case 2, 4, 6:
		fmt.Println("two four or six")
	default:
		fmt.Println("a different number")
	}

	q := 10
	switch {
	case q <= 10:
		fmt.Println("number's less than equal to 10")
		fallthrough
	case q <= 20:
		fmt.Println("number's less than equal to 20")
	default:
		fmt.Println("number's greater than 20")
	}

	typeSwitch()

}
