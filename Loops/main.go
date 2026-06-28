package main

import "fmt"

// only for loop exist in go
// how while loop operate in go, we use for loop to operate while in go.

func main() {

	// while looping using for in Go
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i++
	}

	// Infinite loop in go
	// for {
	// 	fmt.Println("hii")
	// }

	// standard for loop in go
	for k := 1; k <= 10; k++ {
		fmt.Println(k)

		// also use break and continue statement
	}


	// Iterate using range keyword
	for j := range 7 {
      fmt.Println(j)
	}





}
