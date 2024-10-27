package main

import "fmt"
import _ "log"

func main() {

	var name, car string = "Wassem", "Honda Accord 2022"

	// long GO variable declaration
	var locations []string = []string{
		"Canada", "Australia", "Switzerland",
	}

	// Short hand declaration
	school := "University of Toronto"
	profession := "CTO"

	favoriteCars := []string{"BMW", "Merceded"}

	var currency float64 = 12.21

	var (
		shirt    string = "Versace"
		trousers string = "Polo"
	)

	fmt.Println(shirt)
	fmt.Println(trousers)
	footBallClub, player := "Real Madrid", "Ronaldo"

	fmt.Println(footBallClub)
	fmt.Println(player)
	fmt.Println(currency)
	fmt.Println(favoriteCars)

	fmt.Println(profession)
	fmt.Println(school)
	var jobRole = "Creative Technologist"
	fmt.Println(jobRole)
	fmt.Println(locations)
	fmt.Println(name)
	fmt.Println(car)

	fmt.Println("Wassem Darkwa is a creative technologist")

}
