package main

import (
	"fmt"
	"sync"
)

func sayHello() {
	fmt.Println("Hello there")
}

var waitGroup = sync.WaitGroup{}

func main() {

	message := "Hello"
	waitGroup.Add(1)
	go func(message string) {
		fmt.Println(message)
		waitGroup.Done()
	}(message)
	message = "Goodbye"
	waitGroup.Wait()

	// time.Sleep(100 * time.Millisecond)
}
