package main

import "fmt"

/*
Closures: similar to JS
- outer scope variable store inside closure, whenever need this variable it take from closures.


*/

func counter() func() int {
	count := 0

	return func() int {
		count++ // counter store inside closures
		fmt.Println(count)
		return count
	}
}

func main() {
	res := counter()

	res()
	res()
	res()
	res()
	res()
	cnt := res()
	fmt.Println(cnt)
}
