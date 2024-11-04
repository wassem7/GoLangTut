package main

import "fmt"

func main() {

	const age = 23
	const pi float64 = 3.121
	const school, university = "", ""
	// comma separated
	const fullName, profession = "Wassem Darkwa", "Creative Technologist"
	fmt.Println(fullName)
	fmt.Println(profession)
	fmt.Println(age)
	fmt.Println(pi)
	//untyped constants
	const testNumber, testNumber2 = 3, 4

	fmt.Println(testNumber)
	fmt.Println(testNumber2)

	//const declaration block

	const (
		multiplier = 32.1
		divisor    = 12
	)

	// typed constants
	const brand string = "Accord"
	const (
		color = "red"
		car   = "Accord"
	)

	fmt.Println(multiplier)
	fmt.Println(divisor)
}
