package main

import (
	"fmt"
	"runtime"
	"time"
)

/*

Goroutines: A goroutine is a lightweight thread of execution managed entirely by the Go runtime scheduler rather than the operating system kernel.
- used for multithreading and run anything concurrently.

*/

func task(id int) {
	fmt.Println("do task by order", id)
}

func main() {
	cores := runtime.NumCPU()
	fmt.Println(cores)

	for i := 0; i < 10; i++ {
		// go task(i)  // using goroutine we use all working threads.

		go func(id int) {
			fmt.Println("do task by order", id)
		}(i)

	}

	time.Sleep(time.Microsecond * 800)

}
