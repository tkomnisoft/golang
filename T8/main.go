package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// Method 1
func cycleNames(names []string) {
	for _, value := range names {
		sayGreeting(value)
		sayBye(value)
	}
}

// Method 2
func cycleNames2(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("peach")

	// cycleNames2([]string{"mario", "luigi", "peach"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
