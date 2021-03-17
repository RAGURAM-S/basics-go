package main

import (
	"errors"
	"fmt"
)

type greeting struct {
	greeting string
	name     string
}

func printMessage(message string, counter int) {
	fmt.Println(message, counter)
}

func printName(name *string) {
	fmt.Println(*name)
	*name = "Bruce Wayne"
	fmt.Println(*name)
}

// spread operator -> takes all the parameters that are passed into the function and stores them in a slice
func sum(values ...int) int {
	fmt.Printf("values -> %v\n", values)
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func divide(a float64, b float64) (float64, error) {
	if b == 0.0 {
		err := errors.New("math: cannot divide a number by zero")
		return 0, err
	}
	return a / b, nil

}

func main() {
	for i := 1; i <= 5; i++ {
		printMessage("hey there", i)
	}
	name := "Bat Man"
	fmt.Println(name)

	printName(&name)

	fmt.Println(name)

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("the sum is ", total)

	quotient, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("quotient is ", quotient)

	// anonymous functions
	fmt.Println("anonymous functions")
	func() {
		fmt.Println("this is an anonymous function block")
	}() // needs () for the function block to be invoked

	var f func() = func() {
		msg := "function assigned to a variable"
		fmt.Println(msg)
	}

	f()

	g := greeting{
		greeting: "Hello",
		name:     "John Wick",
	}
	// method invokation
	g.greet()
	fmt.Println(g.greeting, g.name)
}

// method declaration
// structs are value types -- to be able to manipulate the struct in a function, it has to be passed as *pointer
func (g *greeting) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Jonathan"
}
