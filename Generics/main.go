package main

import "fmt"

/*

Generics: Prevent duplicate methods implementation.
- In go both works same for defining type -> any
- interface{} == any

*/

// integer slice le liye ye hai or string slice ke liye dusra likhna hoga.
// syntax of generics (any) template
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}

// syntax of generics (comparable) and multiple types template
func printSl[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}

// syntax of generics (scoped) template
func printS[T string | int](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
}

// generics apply on structs
type Stack[T string | int | bool] struct {
	element []T
}

func main() {

	// num := []int{1, 2, 3, 4, 5}
	// lang := []string{"golang", "typescript", "java"}
	//  printS(num)

	stack := Stack[string]{
		element: []string{"golang", "java", "C++"},
	}

	fmt.Println(stack.element)
}
