package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*

Goroutines: Me time sleep use kr rhe the real projects me hume ni pta hota ki kitna time lgega execution me so hum wait groups use krte hai.
- time sleep kya kr rha tha main execution ko stop kr rha tha end hone se.
- ab time sleep ki jgha use krnge wait groups.
- defer: all execution complete hone ke baad run hote hai.
- wait group me only three cheeze hai
    1. wg.add()
		2. wg.done()
		3. wg.wait()

*/

// wait always accept as a pointer
func task(id int, wg *sync.WaitGroup) {

	defer wg.Done() // ye kyu use kiya, bcz jo  humne add kiya tha ab use remove/finish kr rhe hai.

	fmt.Println("do task by order", id)
}

func main() {
	cores := runtime.NumCPU()
	fmt.Println(cores)

	// how to create wait group ->  sync.waitGroup is a type
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1) // why one bcz, only one go routine start.

		// wait group always pass as a reference
		go task(i, &wg) // using goroutine, we use all working threads.

	}


	wg.Wait()

}
