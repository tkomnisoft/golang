package main

import "fmt"

// func main() {
// 	// declaring a string variable
// 	// here we declare a variable nameOne of type string - string can only be double quotes on go
// 	var nameOne string = "John"
// 	fmt.Println("nameOne: ", nameOne)

// 	// here go will infer the type of the variable
// 	var nameTwo = "Luigi"
// 	fmt.Println("nameTwo: ", nameTwo)

// 	var nameThree string
// 	fmt.Println("nameThree: ", nameThree)

// 	// shorthand declaration
// 	// this is the most common way to declare variables in go
// 	// we use the := operator to declare and assign a value to a variable
// 	// we cannot use the := operator outside of a function
// 	nameFour := "Mario"
// 	fmt.Println("nameFour: ", nameFour)

// }

func main() {
	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println("ageOne: ", ageOne, "ageTwo: ", ageTwo, "ageThree: ", ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	// We don't usually specify the size of the int

	// floats
	var scoreOne float32 = 25.98
	var scoreTwo = 42.98
}
