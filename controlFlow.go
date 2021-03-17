package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func callRobots() {

	a := "start"
	defer fmt.Println(a) // this statement prints start
	a = "end"

	const urlString string = "http://www.google.com/robots.txt"

	response, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
	}
	// the statement after defer gets executed at the end of the function. In this case, the response body is closed at the end of the main function
	// defer statements execute in the LIFO order - the last defer statement is executed first the first at last
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", result)
}

func main() {

	callRobots()
}
