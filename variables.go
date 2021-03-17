package main

import (
	"fmt"
	"strconv"
)

// global scope
var Joker string = "Heath Ledger"

// package scope
var (
	ironman         string = "Tony Stark"
	captain_america string = "Steve Rogers"
	god_of_thunder  string = "Thor"
	batman          string = "Bruce Wayne"
)

func main() {

	//  block scope or function scope

	var variable int

	fmt.Printf("value of initialized but unassigned variables - default %v -> %T \n", variable, variable)

	var i int
	i = 32

	var j int = 16

	k := 42

	s := "i'm a string"

	fmt.Printf("value of i -> %v Type %T \n", i, i)
	fmt.Printf("value of j -> %v Type %T \n", j, j)
	fmt.Printf("value of k -> %v Type %T \n", k, k)
	fmt.Printf("value of s -> %v Type %T \n", s, s)

	// type conversion int to float
	temp := float64(i)
	fmt.Printf("temp value %v %T \n", temp, temp)

	s = strconv.Itoa(k)
	fmt.Printf("temp value %v %T", s, s)
}
