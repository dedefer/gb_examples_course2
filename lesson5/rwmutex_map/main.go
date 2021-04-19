package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	globalMap   = map[int]int{}
	globalMapMu = sync.RWMutex{}
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if rand.Float32() < 0.3 {
				globalMapMu.Lock()
				globalMap[i] = i * 100
				globalMapMu.Unlock()
				return
			}
			globalMapMu.RLock()
			fmt.Println(globalMap[i])
			globalMapMu.RUnlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
