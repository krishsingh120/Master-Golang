package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

- Read and write to another file (streaming fashion)

*/

func main() {

	// create a src file where user read data.
	srcFile, err := os.Open("src.txt")
	if err != nil {
		fmt.Println("error during open a file:", err)
		return
	}

	defer srcFile.Close()

	// create a des file
	desFile, err := os.Create("des.txt")
	if err != nil {
		panic(err)
	}

	defer desFile.Close()

	// Streaming setup
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(desFile)

	// infinite loop
	for {
		// reader read a data byte by byte
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		// writer write a data byte by byte
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}

	}

	// Flush remaining or garbage data after loop ends.

	writer.Flush()

	fmt.Println("File Successfully Read and write to another file!")

}
