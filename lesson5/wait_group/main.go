package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i'm goroutine number", i)
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
