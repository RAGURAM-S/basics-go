package main

import "fmt"

func main() {
	var uninitialized bool
	fmt.Printf("%v - %T \n", uninitialized, uninitialized)

	b := true
	fmt.Printf("%v - %T\n", b, b)
	b = false
	fmt.Printf("%v - %T\n", b, b)

	i := 10             // 1010
	j := 3              // 0011
	fmt.Println(i & j)  // 0010 - 2
	fmt.Println(i | j)  // 1011 - 11
	fmt.Println(i ^ j)  // 1001 - 9
	fmt.Println(i &^ j) //

	x := 8 // 1000

	fmt.Println(x << 3) // left shifting  --> 0000 1000 --> 01000000  <3 zeros shifted from the left to the right>
	fmt.Println(x >> 3) // right shifting --> 0000 1000 --> 00000001  <3 zeros shifted from the right to the left>

	// float

	p := 25
	q := 10.5
	fmt.Println(float64(p) * q)

	// strings --> utf-8

	str := "Hello there"
	fmt.Printf("%v - %T\n", str, str)

	fmt.Printf("%v - %T\n", str[0], str[0])

	fmt.Printf("%v - %T\n", string(str[0]), string(str[0]))

	bytes := []byte(str)
	fmt.Printf("%v - %T\n", bytes, bytes)

	// runes --> utf-32

	var t rune = 'a'
	fmt.Printf("%v - %T\n", t, t)

}
