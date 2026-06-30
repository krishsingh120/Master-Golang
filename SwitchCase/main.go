package main

import (
	"fmt"
)

// no need a break statement inside switch bcz they will automatically handle.

func main() {
	// simple switch

	// i := 1
	// switch i {
	// case 1:
	// 	fmt.Println("1")
	// case 2:
	// 	fmt.Println("2")
	// case 3:
	// 	fmt.Println("3")
	// default:
	// 	fmt.Println("Infinity")
	// }

	// // multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Working day")

	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("this is integer", t)
		case string:
			fmt.Println("this is string", t)
		case bool:
			fmt.Println("this is boolean", t)
		}
	}

	whoAmI("12")

}
