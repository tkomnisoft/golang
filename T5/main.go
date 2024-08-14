package main

import "fmt"

func main() {
	// Go has fixed array sizes
	// arrays
	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages)

	var ages1 = [3]int{20, 25, 30}
	fmt.Println(ages1, len(ages1))

	ages2 := [3]int{20, 25, 30}
	fmt.Println(ages2)

	// slices

	// slices are like arrays but they are dynamic
	// slices don't have a fixed size
	var scores = []int{100, 50, 60}

	// replaces the value at index 2 with 25
	scores[2] = 25

	// adds 85 to the end of the slice
	// we can use append for slices
	scores = append(scores, 85)

	// print scores
	fmt.Println(scores, len(scores))
}
