package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	mu = sync.RWMutex{}
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if rand.Float32() < 0.3 {
				mu.Lock()

				fmt.Printf("goroutine %d writing 1\n", i)
				fmt.Printf("goroutine %d writing 2\n", i)

				mu.Unlock()
				return
			}
			mu.RLock()

			fmt.Printf("goroutine %d reading 1\n", i)
			fmt.Printf("goroutine %d reading 2\n", i)

			mu.RUnlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
