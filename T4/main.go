package main

import "fmt"

func main() {
	// Print and Println
	// fmt.Print("Hello, ")
	// fmt.Println("World!")
	age := 20
	name := "John"

	// print variables
	fmt.Println("my age is", age, "and my name is", name)

	// Use quote to print the value of the variable
	fmt.Printf("my age is %q and my name is %q \n", 11, name)

	// Print the type of the variable
	fmt.Printf("age is of type %T \n", age)

	// print float with 2 decimal points. the f stands for float
	fmt.Printf("you scored %0.2f points! \n", 355.55)
}

// https://pkg.go.dev/fmt
