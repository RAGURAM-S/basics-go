package main

import (
	"bytes"
	"fmt"
)

func main() {
	nameString := "Hello there"
	var writer Writer = ConsoleWriter{}
	writer.Write([]byte(nameString))

	value := IntCounter(0)
	var inc Incrementer = &value
	fmt.Println(inc.Increment())

	var printerCloser PrinterCloser = NewBufferedPrinterCloser()
	printerCloser.Print([]byte("Hello YouTube Listeners, this is a test"))
	printerCloser.Close()

	r, ok := printerCloser.(*BufferedPrinterCloser)
	if ok {
		// fmt.Println(ok)
		fmt.Println(r)
	} else {
		fmt.Println("Error in Conversion")
	}

	fmt.Println()

	// empty interface
	var emptyInterface interface{} = NewBufferedPrinterCloser()
	if printerCloser, ok = emptyInterface.(PrinterCloser); ok {
		printerCloser.Print([]byte("Hello YouTube Listeners, this is a test"))
		printerCloser.Close()
	}

	r, ok = printerCloser.(*BufferedPrinterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Error in Conversion")
	}

	//
	var x interface{} = 0

	switch x.(type) {
	case int:
		fmt.Println("an integer")
	case string:
		fmt.Println("a string")
	default:
		fmt.Println("unknown data type")
	}

}

type Writer interface {
	// write method which accepts a slice of bytes and returns an int and an error
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	fmt.Println(n, err)
	return n, err
}

type IntCounter int

type Incrementer interface {
	Increment() int
}

func (number *IntCounter) Increment() int {
	*number++
	return int(*number)
}

type Printer interface {
	Print([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type PrinterCloser interface {
	Printer
	Closer
}

type BufferedPrinterCloser struct {
	buffer *bytes.Buffer
}

func (bpc *BufferedPrinterCloser) Print(data []byte) (int, error) {
	n, err := bpc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bpc.buffer.Len() >= 8 {
		_, err := bpc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		// fmt.Printf("%v - %T\n", string(v), v)

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bpc *BufferedPrinterCloser) Close() error {
	for bpc.buffer.Len() > 0 {
		data := bpc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedPrinterCloser() *BufferedPrinterCloser {
	return &BufferedPrinterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
