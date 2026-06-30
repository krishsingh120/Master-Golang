package main

import (
	"fmt"
	"maps"
)

/*
maps: hash, objects, dist

IMP: if key does not exist in map, they will return zero value.
*/
func main() {
	// create a map
	var freq = map[string]int{"a": 2, "b": 1, "c": 4}
	fmt.Println(freq["a"])

	freq1 := map[string]string{"a": "2", "b": "1", "c": "4"}
	fmt.Println(freq1["c"])

	// var map1 map[string]int
	// map1["A"] = 100  // Error -> panic: assignment to entry in nil map
	// fmt.Println(map1)

	// using make
	map1 := make(map[string]int)
	fmt.Println(map1) // this is the sol of panic error.
	map1["A"] = 23
	map1["B"] = 3
	map1["D"] = 54
	map1["G"] = 21

	fmt.Println(map1)
	fmt.Println(len(map1))

	delete(map1, "A")
	fmt.Println(map1)

	// clear(map1)
	// fmt.Println(map1)

	val, exists := map1["G"]

	if exists {
		fmt.Println("Exist: ", val)
	} else {
		fmt.Println("not Exist")
	}

	m1 := map[string]int{"a": 2, "b": 1, "c": 4}
	m2 := map[string]int{"a": 2, "b": 1, "c": 4}

	fmt.Println(maps.Equal(m1, m2))

}
