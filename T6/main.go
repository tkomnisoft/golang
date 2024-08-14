// Standard Package
package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello")) // This function is case sensitive
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
}
