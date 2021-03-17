package main

import "fmt"

type SuperHero struct {
	number        int
	name          string
	superHeroName string
	companions    []string
}

func main() {

	ironMan := SuperHero{
		number:        1,
		name:          "Tony Stark",
		superHeroName: "Iron Man",
		companions: []string{
			"Pepper Potts",
			"Steve Rogers",
			"Natasha Romanoff",
			"Bruce Banner",
			"Thor",
			"Happy Hogan",
			"Peter Parker",
		},
	}

	batMan := SuperHero{
		number:        2,
		name:          "Bruce Wayne",
		superHeroName: "Bat Man",
		companions: []string{
			"Alfred Pennyworth",
			"James Gordon",
			"Rachel Dawes",
			"Robin",
			"Joker",
			"Lucius Fox",
		},
	}

	// anonymous struct

	doctorStrange := struct {
		number        int
		name          string
		superHeroName string
		companions    []string
	}{
		number:        3,
		name:          "Doctor Strange",
		superHeroName: "Doctor Strange",
		companions: []string{
			"Wong",
			"The Elder One",
		},
	}

	fmt.Printf("%v - %v\n", ironMan.superHeroName, ironMan)

	fmt.Printf("%v - %v\n", batMan.superHeroName, batMan)

	fmt.Printf("%v - %v\n", doctorStrange.superHeroName, doctorStrange)

	anotherBatMan := batMan
	anotherBatMan.superHeroName = "The Dark Knight"

	fmt.Printf("batMan - %v\n", batMan.superHeroName)
	fmt.Printf("anotherBatMan - %v\n", anotherBatMan.superHeroName)

	// reference type
	nextBatMan := &batMan
	nextBatMan.superHeroName = "The Dark Knight"

	fmt.Printf("batMan - %v\n", batMan.superHeroName)
	fmt.Printf("anotherBatMan - %v\n", nextBatMan.superHeroName)

}
