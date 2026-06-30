package main

import (
	"fmt"
)

// Array: number of sequence with specific length

/*
Default value of array:
	int , float :0
	string :"" empty string
	bool : false
	pointer/interface :nil


	when we use Arrays in go:
	   - fixed size, that is predictable
		 - memory optimization
		 - constant time access

*/

func main() {
	// Syntax

	// 1. initialize empty fixed size array.
	var num [3]int
	fmt.Println(len(num))

	// 2. initialize fixed size array with some values.
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 3. initialize variable size array with some values.
	var brr = [...]string{"apple", "grapes", "cherry", "pineapple"}
	fmt.Println(brr)

	// shorthand
	languages := [...]string{"Javascript", "java", "Go", "python"}
	fmt.Println(languages)

	// Partial initialization of an integer array
	arr1 := [5]int{10, 20}
	fmt.Println(arr1)

	arr2 := [5]int{1: 10, 2: 40} // {index number: value}
	fmt.Println(arr2)            // Output: [0 10 40 0 0]
	// Here, 10 is assigned to index 1 and 40 to index 2.
	// The rest of the elements (index number 0, 3, 4) default to 0.

	// 2-D array in go
	grid := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(grid)
}
