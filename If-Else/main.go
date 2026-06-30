package main

import "fmt"

func main() {
	age := 16

	if age < 18 {
		fmt.Println("Person is not an adult")
	} else if age == 18 {
		fmt.Println("Person is just eligible to adult works in life")
	} else {
		fmt.Println("Person is proper eligible to adult works in life")
	}

	// also possible variable initialize inside if condition
	if j := 12; j <= 20 {
		fmt.Println("Win!")
	} else {
		fmt.Println("Loss!")
	}
}
