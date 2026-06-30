package main

import (
	"fmt"
	"slices"
)

/*
Slices:
1. dynamic size
2. most used constructor in Go.
3. + useful methods.
4. In go nil == null in other language.


cap: Maximum no. of element can fit (dynamic it can be change).



*/

func main() {
	// syntax of slice

	// declare but not initialized.
	var num []int // nil value
	fmt.Println(len(num))
	fmt.Println(num == nil)

	/*
		make -> user decide initial cap or len of slice.
		syntax -> make(type, len, cap)
	*/

	var arr = make([]int, 2, 5)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	fmt.Println(arr)
	/*
		    arr look like this = [0, 0, 1, 2]
				starting two value represent slice, bcz we use a make method and initialize with size 2.
				standard way to use make method len of slice is 0.
				if we reach initialize cap to cap by default double the own size.
				initial -> 5 -> 10 -> 20....
	*/
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	brr := []int{}
	brr = append(brr, 1)
	brr = append(brr, 2)
	fmt.Println(brr)

	// copy function
	arr1 := make([]int, 0, 5)
	arr1 = append(arr1, 3)
	arr1 = append(arr1, 8)

	arr2 := make([]int, len(arr1))

	fmt.Println(arr1, arr2)

	copy(arr2, arr1)
	fmt.Println(arr1, arr2)

	// slice operator
	var nums = []int{4, 3, 5, 1, 2}

	// extract element from start to end(excluded)
	fmt.Println(nums[2:4]) // 2 to 4
	fmt.Println(nums[2:]) // 2 to end
	fmt.Println(nums[:4]) // 0 to 4


	// slice package inside this many more methods
	fmt.Println(slices.Equal(arr1, arr2)) // return bool


}
