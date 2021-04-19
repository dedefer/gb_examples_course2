package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	globalMap = sync.Map{}
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if rand.Float32() < 0.3 {
				// globalMap[i] = i*100
				globalMap.Store(i, i*100)
				return
			}

			valI, _ := globalMap.Load(i)
			val, _ := valI.(int)
			fmt.Printf("%+v\n", val)
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
