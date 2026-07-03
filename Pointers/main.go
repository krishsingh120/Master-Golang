package main

import "fmt"

/*
Pointers: similar to c/c++.

- we store variable/data structure in memory location, so pointer basically is the address of that memory location.
- so if i want to change a value of variable permanently, so pass value = as reference using pointers.
- means not pass variable copy, pass real memory location address, that's why change happen permanently.

- how to access address of memory location of any variable -> like this = &<variable_name>

*/

// num : pass as a = by value
// means: in this func we use num as a copy not a real value so that's why in main func num not changed.
func changeNum(num int) {
	num = 5

	fmt.Println("after change num value in changeNum func:", num)
}

// change val by reference
func changeVal(val *int) {
	*val = 5 // change krne ke liye * use krna hoga, btana hota hai ki address location pe new value assign kr rhe hai.

	fmt.Println("after change num value in changeNum func:", val)

}

func main() {

	num := 1

	// fmt.Println("num in main fun before changeNum call:", num)

	// changeNum(num)

	// fmt.Println("num in main fun after changeNum call:", num)

	fmt.Println("Memory location of this num:", &num)
	fmt.Println("num in main fun before changeVal call:", num)

	changeVal(&num)

	fmt.Println("num in main fun after changeVal call:", num)

}
