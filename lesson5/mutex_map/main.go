package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	globalMap   = map[int]int{}
	globalMapMu = sync.Mutex{}
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			globalMapMu.Lock()
			defer globalMapMu.Unlock()

			if rand.Float32() < 0.3 {
				globalMap[i] = i * 100
				return
			}
			fmt.Println(globalMap[i])
		}(i)
	}

	wg.Wait()

	fmt.Println("main done", len(globalMap))
}
