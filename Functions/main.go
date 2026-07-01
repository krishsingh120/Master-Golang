package main

import "fmt"

/*
functions: func in go they supports to return multiple values.
1. func in go are first class functions.
2. bcz func pass as a argument, also accept as parameter, also function return a another function.

for backend projects syntax:

func server() (data, error) {
  return data, error
	}

*/

// func hello(a int, b int) equal to -> hello(a, b int)
func add(a int, b int) int {
	return a + b
}

func getLanguage() (string, string, string) {
	return "golang", "javascript", "java"
}

// func returns another func
func processIt() func(a int) int {
	cnt := 2
	return func(val int) int {
		fmt.Println(val + cnt)
		return cnt
	}
}

func main() {
	// res := add(12, 4)
	// fmt.Println(res)

	// l1, l2, l3 := getLanguage()

	// fmt.Println(l1, l2, l3)
	// fmt.Println(getLanguage())

	fn := processIt()
	fmt.Println(fn(12))

}
