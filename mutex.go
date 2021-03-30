package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello() {
	// m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	waitGroup.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	waitGroup.Done()
}

func main() {
	fmt.Printf("Number of threades - %v\n", runtime.GOMAXPROCS(-1))

	// runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()

	}
	waitGroup.Wait()

}

//checks for race condition in a multi-threaded application => go run -race mutex.go
