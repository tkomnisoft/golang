package main

import (
	"fmt"
)

func main() {
	// 1. One way to loop in go
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// 2. Another way to loop in go
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is:", i)
	// 	// Will print from 0 to 4
	// }

	// 3. Slice printing in a loop
	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
	// fmt.Println("Length size: ", len(names))
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(i, names[i])
	// }

	// 4. Range keyword
	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}
	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	// 5. If we don't want to use the index
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

}
