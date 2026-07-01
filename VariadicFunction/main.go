package main

import "fmt"

/*

Variadic function: those func accepts a any no. of parameters.
fmt.Println(1, 2, 3, "hello", true)

any type in go -> any ==> interface{}


*/

func sum(nums ...int) int {
	total := 0

	for _, val := range nums {
		total += val
	}

	return total
}

func main() {
	// println -> variadic func
	fmt.Println(1, 2, "hello", true)

	// res := sum(1, 2, 3, 4)
	// fmt.Println(res)

	nums := []int{1, 2, 3, 4, 5}
	ans := sum(nums...)
	fmt.Println(ans)

}
