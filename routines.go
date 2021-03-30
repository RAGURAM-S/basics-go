package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}
var counter = 0

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	waitGroup.Done()
}

func increment() {
	counter++
	waitGroup.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		go sayHello()
		go increment()
	}
	waitGroup.Wait()

}
