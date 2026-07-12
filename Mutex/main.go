package main

import (
	"fmt"
	"sync"
)

/*

Mutex: To prevent Race condition we use mutex.
- Mutex = Mutual Exclusion. It's a lock. Only one goroutine can hold the lock at a time. Everyone else must wait their turn.

*/

type Post struct {
	view int
	mu   sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	defer p.mu.Unlock() // taaki kuch bhi ho ye code hamesha chle.

	p.mu.Lock() // lock sirf utna kro jitna hume lock krna ho - Means just uske upper lock lga jha need ho.
	p.view += 1
}

func main() {
	// fmt.Println("hello")

	var wg sync.WaitGroup
	newPost := Post{view: 0}

	// Race condition with multiple goroutines.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go newPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(newPost.view) // without Mutex this maybe 100, or not be.
}
