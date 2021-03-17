package main

import (
	"fmt"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	CanFly   bool
	SpeedKph float32
}

func main() {
	b := Bird{}
	b.Name = "Cassovary"
	b.Origin = "Africa"
	b.CanFly = false
	b.SpeedKph = 45
	fmt.Println(b)

	a := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		CanFly:   false,
		SpeedKph: 48,
	}

	fmt.Println(a)

	// 	types := reflect.TypeOf(Animal{})
	// 	field = types.FieldByName("Name")
	// 	fmt.Println(field.Tag)
}
