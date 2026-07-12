package main

import (
	"fmt"
	"os"
)

/*

File Handling: In go there are many ways, but i learn core standard way using OS Modules.

- Using this OS module we can read a files and folders both of them.

*/

func main() {
	fmt.Println("Hi!")

	// // create a File using go OS module
	// file, err := os.Create("data.txt")

	// if err != nil {
	// 	fmt.Println("Error during file creation error:", err)
	// 	return
	// }

	// defer file.Close()

	// // writing a file in go using OS modules.
	// data := "Hello this is go Data file..."
	// _, err = file.WriteString(data)
	// if err != nil {
	// 	fmt.Println("Error during writing a file:", err)
	// 	return
	// }

	// // Read a file in go using OS modules.

	// // file, err = os.Open("data.txt")
	// // if err != nil {
	// // 	fmt.Println("Error during opening a file:", err)
	// // 	return
	// // }

	// // defer file.Close()

	// // // create a buffer to store and read a file in form of buffer slice.
	// // buffer := make([]byte, 1024)
	// // bytesRead, err := file.Read(buffer)
	// // if err != nil {
	// // 	fmt.Println("Error reading from file:", err)
	// // 	return
	// // }

	// // fmt.Println("Read:", string(buffer[:bytesRead])) // Output: Read: Hello, World!

	// // Another way to read a file - when use if file is small.
	// content, err := os.ReadFile("data.txt")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println("Read:", string(content))

	// /// --------------------------------- ///
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat() // all info related to this file (f)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())           // File ka naam (sirf naam, path nahi) — e.g. data.txt
	// fmt.Println("file or folder:", fileInfo.IsDir())     // true agar folder hai, false agar file hai
	// fmt.Println("file size:", fileInfo.Size())           // File size bytes mein (int64)
	// fmt.Println("file permission:", fileInfo.Mode())     // File permissions + type (e.g. -rw-r--r--)
	// fmt.Println("file modified at:", fileInfo.ModTime()) // Last modified timestamp (time.Time)



	// //// --------------------------- ////
	// // Read both Folders and Files.

	// // dir, err := os.Open("./") // curr dir
	// dir, err := os.Open("../") // One dir back from curr dir
	// if err != nil {
	// 	fmt.Println("Error during Reading folders and files:", err)
	// 	return
	// }

	// defer dir.Close()

	// dirInfo, err := dir.ReadDir(-1)

	// for _, info := range dirInfo {
	// 	fmt.Println(info.Name(), info.IsDir())
	// }


	//// ------------- ////
	// Delete a file in go

	err := os.Remove("del.txt")
	if err != nil {
		panic(err)
	}
	
















}
