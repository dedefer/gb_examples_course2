package main

import (
	"sync"
)

func main() {
	arr := map[int]int{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			arr[i%10] = i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			arr[i%10] = i + 1
		}
	}()

	wg.Wait()
}
