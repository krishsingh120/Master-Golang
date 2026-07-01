package main

import "fmt"

/*

Range: use for iterating over Data structure


*/
func main() {
	// iterate over slice
	num := []int{1, 2, 3, 4, 5, 6}
	for idx, val := range num {
		fmt.Println("index", idx, "->", "value", val)
	}

	// iterate over map
	m := map[string]string{"name": "krish", "age": "22", "profession": "developer", "location": "bareilly"}

	for k, v := range m {
		fmt.Println("key", k, "->", "value", v)
	}

	// on strings -> in go character data type(char) are not present, but in go runes -> they basically return unicode of characters.
	for i, ch := range "golang" {
		fmt.Println(i, "-> ", string(ch), "unicode of this character", ch)
	}
}
