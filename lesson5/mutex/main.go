package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	mu = sync.Mutex{}
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			if rand.Float32() < 0.3 {
				fmt.Printf("goroutine %d writing 1\n", i)
				fmt.Printf("goroutine %d writing 2\n", i)
				return
			}

			fmt.Printf("goroutine %d reading 1\n", i)
			fmt.Printf("goroutine %d reading 2\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
