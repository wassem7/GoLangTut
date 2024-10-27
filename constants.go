package main

import "fmt"

func main() {

	const age = 23
	const pi float64 = 3.121
	// comma separated
	//typed constants
	const fullName, profession = "Wassem Darkwa", "Creative Technologist"
	fmt.Println(fullName)
	fmt.Println(profession)
	fmt.Println(age)
	fmt.Println(pi)
	//untyped constants
	const testNumber, testNumber2 = 3, 4

	fmt.Println(testNumber)
	fmt.Println(testNumber2)
}
