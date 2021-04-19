package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	pool := sync.Pool{}
	pool.New = func() interface{} { return &strings.Builder{} }

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			b := pool.Get().(*strings.Builder)

			b.WriteString("i'm goroutine number")
			b.WriteString("i'm goroutine number 2")
			b.WriteString("i'm goroutine number 3")

			fmt.Println(b.String())

			b.Reset()
			pool.Put(b)
		}(i)
	}

	wg.Wait()

	fmt.Println("main done")
}
