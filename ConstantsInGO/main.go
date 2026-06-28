package main

import "fmt"

const height = 171

func main() {
	const name = "krish"
	fmt.Println(name)

	const age = 22
	fmt.Println(age)

	fmt.Println(height)

	// multiple const grouping
	const (
		port = 3000
		host = 8001
	)

	fmt.Println(port, host)

	// name = "singh"
}
