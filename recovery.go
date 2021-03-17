package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("function about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// if the panic cannot be handled, a panic is to be added to the anonymous function block, thereby ending the program
			panic(err)
		}
	}()
	panic("An unkonwn error occurred!!")
	fmt.Println("done panicking")
}

func main() {
	fmt.Println("app starts")
	panicker()
	fmt.Println("app ends")
}
