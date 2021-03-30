package main

import (
	"fmt"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	waitGroup.Add(2)
	// receiving data from the channel
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		waitGroup.Done()
	}()
	// sending data to the channel
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
