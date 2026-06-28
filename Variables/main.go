package main

import "fmt"

// using var keywords we can declare a variable outside the main function
var lang = "golang"

// outside shorthand does not work.

func main() {
	// way-1 to create variables in go
	var name string = "go dev"
	fmt.Println(name)

	// way-2 to create variables in go
	var fullName = "krish singh"
	fmt.Println(fullName)

	var age int = 22
	fmt.Println("Age of", fullName, "is:", age)

	// Shorthand to declare a variable and assign a value
	height := 171
	fmt.Println(height)

	fmt.Println(lang)
}
