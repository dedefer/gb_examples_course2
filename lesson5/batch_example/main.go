package main

import (
	"fmt"
	"sync"
)

func main() {
	batches := [][]int{}
	batch := make([]int, 0, 16)
	for i := 0; i < 1000; i++ {
		batch = append(batch, i)

		if len(batch) == 16 {
			batches = append(batches, batch)
			batch = make([]int, 0, 16)
		}
	}

	if len(batch) != 0 {
		batches = append(batches, batch)
	}

	wg := sync.WaitGroup{}

	for i, batch := range batches {
		wg.Add(1)
		go func(i int, batch []int) {
			defer wg.Done()
			fmt.Printf("processing batch %d %+v\n", i, batch)
		}(i, batch)
	}

	wg.Wait()
}
